# Screen 6: Export a new CV

- Status: Draft; wording to iron out
- Session: [mvp-session.md](../mvp-session.md), step 4; story [US-08](../user-stories/US-08-new-cv.md)
- Found by playing Anne's first visit, the interface answering as a person

Four steps: who, what, look, file. The result is saved as a CV and can be built again with the facts added since.

## ① Who

```
New CV                          ① Who  ② What  ③ Look  ④ File

Who is this CV from?

  (•) Anne Martin · anne@… · Lyon
  ( ) Add another identity
                                                        [Next]
```

## ② What

The view editor, with a call to action to go on without filtering.

```
Export your CV                                          [Next]

On this CV: 4 occupations, 13 practices.

Occupations   ☑ Backend developer, ACME Logistics     2021 – today
              ☑ Maintainer of mqttlite                2019 – today
              ☑ Embedded software engineer, Nordwind  2017 – 2021
              ☑ MSc Computer Science, Rhône Univ.     2015 – 2017
Period        from [    ]  to [    ]
Show          ☑ levels   ☐ excerpts from my old CV

Practices     ☑ Develop data processing applications · Kafka pipeline · Rust   2021 – today
              ☑ …
```

## ③ Look

```
New CV                          ① Who  ② What  ③ Look  ④ File

Pick a look.

  Installed    [preview] Even   [preview] Flat
  More         Elegant · Kendall · …    [Install]
                                                        [Next]
```

A preview shows Anne's CV in that look, rendered by the `cv` plugin: [Use this look] [Back to looks].

## ④ File

```
New CV                          ① Who  ② What  ③ Look  ④ File

Save as     [CV #1                    ]
File        (•) PDF   ( ) HTML
                                               [Build and download]
```

## Build

```
CV #1                          ① Who  ② What  ③ Look  ④ File

Building the PDF…   ✓ Facts selected   ✓ CV written   ● Rendered   ○ Downloaded
                                                          [Stop]
```

then:

```
CV #1 is ready: anne-martin-cv-1.pdf          [Download again]

Saved in your CVs. Build it again later to include new facts.

  [Open my CVs]   [Home]
```

## Rules found

- Who comes before what: an identity can tie a CV to some activities.
- The identity, view, look and file are saved together as a CV, in the owner's preferences.
- Building is a task with stages: facts selected, CV written (core), rendered (`cv` plugin), downloaded.

## Open

- How the view drives the filter, and how to focus a CV on a job.
- Preferences: where saved CVs are kept (specification).
- "My CVs" screen.
- Themes: which are installed, where more are installed from.
