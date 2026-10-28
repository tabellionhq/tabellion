// Command sqlite checks what ADR-0022 requires of ncruces/go-sqlite3 before relying on it:
// encryption with a service key and with a person's passphrase, full-text search,
// vector search, start time and memory. It is a throwaway measurement, not product code.
//
// Usage:
//
//	sqlite checks <dir>        run every check, print YAML on stdout
//	sqlite start <db> <hexkey> open an encrypted database and read it, then exit
package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	mrand "math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ncruces/go-sqlite3"
	"github.com/ncruces/go-sqlite3/driver"
	"github.com/ncruces/go-sqlite3/ext/fts5"
	"github.com/ncruces/go-sqlite3/ext/vec1"
	_ "github.com/ncruces/go-sqlite3/vfs/adiantum"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

const secret = "Anne Martin, backend developer at ACME Logistics"

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: sqlite checks <dir> | sqlite start <db> <hexkey>")
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "checks":
		err = checks(os.Args[2])
	case "start":
		if len(os.Args) < 4 {
			err = errors.New("start needs <db> <hexkey>")
			break
		}
		err = start(os.Args[2], os.Args[3])
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func open(dsn string) (*sql.DB, error) {
	return driver.Open(dsn, func(c *sqlite3.Conn) error {
		if err := fts5.Register(c); err != nil {
			return err
		}
		return vec1.Register(c)
	})
}

func encrypted(path, param, key string) string {
	return fmt.Sprintf("file:%s?vfs=adiantum&%s=%s", path, param, key)
}

func checks(dir string) error {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	ctx := context.Background()
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return err
	}
	hexkey := hex.EncodeToString(key)

	fmt.Println("encryption_service_key:")
	if err := checkEncryption(ctx, filepath.Join(dir, "service.db"), "hexkey", hexkey, "hexkey", strings.Repeat("0", 64)); err != nil {
		return fmt.Errorf("service key: %w", err)
	}
	fmt.Println("encryption_person_passphrase:")
	if err := checkEncryption(ctx, filepath.Join(dir, "person.db"), "textkey", "correct horse battery staple", "textkey", "wrong horse"); err != nil {
		return fmt.Errorf("passphrase: %w", err)
	}
	fmt.Println("full_text:")
	if err := checkFullText(ctx, filepath.Join(dir, "labels.db"), hexkey); err != nil {
		return fmt.Errorf("full text: %w", err)
	}
	fmt.Println("vectors:")
	if err := checkVectors(ctx); err != nil {
		return fmt.Errorf("vectors: %w", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "hexkey"), []byte(hexkey), 0o600); err != nil {
		return err
	}
	fmt.Printf("memory:\n  peak_rss_mib: %s\n", peakRSS())
	return nil
}

// checkEncryption writes a record, then checks the raw file hides it, the right key reads it back,
// and a wrong key is refused.
func checkEncryption(ctx context.Context, path, param, key, wrongParam, wrongKey string) error {
	os.Remove(path)
	t := time.Now()
	db, err := open(encrypted(path, param, key))
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, `CREATE TABLE fact(id INTEGER PRIMARY KEY, label TEXT)`); err != nil {
		db.Close()
		return err
	}
	for i := range 1000 {
		if _, err := db.ExecContext(ctx, `INSERT INTO fact(label) VALUES(?)`, fmt.Sprintf("%s #%d", secret, i)); err != nil {
			db.Close()
			return err
		}
	}
	db.Close()
	write := time.Since(t)

	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	hidden := !bytes.Contains(raw, []byte("Anne Martin")) && !bytes.HasPrefix(raw, []byte("SQLite format 3"))

	t = time.Now()
	db, err = open(encrypted(path, param, key))
	if err != nil {
		return err
	}
	var n int
	err = db.QueryRowContext(ctx, `SELECT count(*) FROM fact WHERE label LIKE 'Anne%'`).Scan(&n)
	db.Close()
	if err != nil {
		return fmt.Errorf("read with the right key: %w", err)
	}
	read := time.Since(t)

	db, err = open(encrypted(path, wrongParam, wrongKey))
	if err != nil {
		return err
	}
	var m int
	wrongErr := db.QueryRowContext(ctx, `SELECT count(*) FROM fact`).Scan(&m)
	db.Close()

	fmt.Printf("  plaintext_absent_from_file: %t\n", hidden)
	fmt.Printf("  rows_read_with_right_key: %d\n", n)
	fmt.Printf("  wrong_key_refused: %t\n", wrongErr != nil)
	if wrongErr != nil {
		fmt.Printf("  wrong_key_error: %q\n", wrongErr.Error())
	}
	fmt.Printf("  write_1000_rows_ms: %d\n", write.Milliseconds())
	fmt.Printf("  open_and_read_ms: %d\n", read.Milliseconds())
	if !hidden || n != 1000 || wrongErr == nil {
		return errors.New("check failed")
	}
	return nil
}

// checkFullText indexes labels at the size of ESCO (about 13,900 skills, in 3 languages)
// in an encrypted database, then searches a word.
func checkFullText(ctx context.Context, path, hexkey string) error {
	os.Remove(path)
	db, err := open(encrypted(path, "hexkey", hexkey))
	if err != nil {
		return err
	}
	defer db.Close()
	if _, err := db.ExecContext(ctx, `PRAGMA temp_store = memory; CREATE VIRTUAL TABLE label USING fts5(text, lang UNINDEXED)`); err != nil {
		return err
	}
	words := []string{"develop", "software", "manage", "data", "pipeline", "design", "test", "teach", "analyse", "build", "operate", "maintain"}
	r := mrand.New(mrand.NewPCG(1, 2))
	t := time.Now()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for i := range 13900 * 3 {
		text := fmt.Sprintf("%s %s %s skill %d", words[r.IntN(len(words))], words[r.IntN(len(words))], words[r.IntN(len(words))], i)
		if _, err := tx.ExecContext(ctx, `INSERT INTO label(text, lang) VALUES(?, ?)`, text, []string{"en", "fr", "de"}[i%3]); err != nil {
			tx.Rollback()
			return err
		}
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO label(text, lang) VALUES('use software design patterns', 'en')`); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	load := time.Since(t)
	t = time.Now()
	var found string
	err = db.QueryRowContext(ctx, `SELECT text FROM label WHERE label MATCH 'patterns' LIMIT 1`).Scan(&found)
	if err != nil {
		return err
	}
	fmt.Printf("  labels_indexed: %d\n", 13900*3+1)
	fmt.Printf("  index_ms: %d\n", load.Milliseconds())
	fmt.Printf("  search_ms: %d\n", time.Since(t).Milliseconds())
	fmt.Printf("  found: %q\n", found)
	if found != "use software design patterns" {
		return errors.New("check failed")
	}
	return nil
}

// checkVectors stores made-up vectors and checks the nearest neighbour of a query is the expected one.
// It shows the search works; it says nothing of the quality of real embeddings.
func checkVectors(ctx context.Context) error {
	db, err := open("file:/vec.db?vfs=memdb")
	if err != nil {
		return err
	}
	defer db.Close()
	const dim, count = 384, 20000
	if _, err := db.ExecContext(ctx, `
		CREATE VIRTUAL TABLE concept USING vec1;
		INSERT INTO concept(cmd, arg) VALUES('rebuild', '{index:"flat", distance:"cos"}');`); err != nil {
		return err
	}
	r := mrand.New(mrand.NewPCG(3, 4))
	vectors := make([][]float32, count)
	t := time.Now()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for i := range count {
		v := make([]float32, dim)
		for j := range v {
			v[j] = float32(r.NormFloat64())
		}
		vectors[i] = v
		if _, err := tx.ExecContext(ctx, `INSERT INTO concept(rowid, vector) VALUES(?, ?)`, i+1, blob(v)); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	load := time.Since(t)

	// The query is concept 4242 with a little noise: it must come first.
	q := make([]float32, dim)
	for j := range q {
		q[j] = vectors[4241][j] + float32(r.NormFloat64()*0.05)
	}
	t = time.Now()
	rows, err := db.QueryContext(ctx, `SELECT rowid FROM concept(?, '{K: 5}')`, blob(q))
	if err != nil {
		return err
	}
	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			rows.Close()
			return err
		}
		ids = append(ids, id)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Printf("  vectors: %d\n  dimensions: %d\n  mode: flat index, cosine, exact search\n", count, dim)
	fmt.Printf("  load_ms: %d\n", load.Milliseconds())
	fmt.Printf("  query_ms: %d\n", time.Since(t).Milliseconds())
	fmt.Printf("  nearest: %v\n", ids)
	if len(ids) == 0 || ids[0] != 4242 {
		return errors.New("check failed")
	}
	return nil
}

func blob(v []float32) []byte {
	b := make([]byte, 4*len(v))
	for i, f := range v {
		binary.NativeEndian.PutUint32(b[4*i:], math.Float32bits(f))
	}
	return b
}

// start is what a server does when it wakes: open the encrypted record and read it.
func start(path, hexkey string) error {
	db, err := open(encrypted(path, "hexkey", hexkey))
	if err != nil {
		return err
	}
	defer db.Close()
	var n int
	if err := db.QueryRow(`SELECT count(*) FROM fact`).Scan(&n); err != nil {
		return err
	}
	fmt.Printf("rows: %d\npeak_rss_mib: %s\n", n, peakRSS())
	return nil
}
