# libmusic

[![Build Status](https://travis-ci.org/moul/libmusic.svg?branch=master)](https://travis-ci.org/moul/libmusic)
[![GoDoc](https://godoc.org/github.com/moul/libmusic?status.svg)](https://godoc.org/github.com/moul/libmusic) [![GuardRails badge](https://badges.production.guardrails.io/moul/libmusic.svg)](https://www.guardrails.io)

Manipulate Music in Golang

## Status

Currently writing foundations (music theory, music physics), see [godoc](https://godoc.org/github.com/moul/libmusic).

Next steps will be to develop real programs and examples.

## Features

The following list is a mix between _things that are already done_ and _new ideas_, probably not up-to-date during the dev phase.

* Note
  * [ ] Parse from string
  * [ ] Interval operations
  * [ ] Comparisons
  * [ ] Exported as go symbols
  * [ ] Semitone
  * [ ] Cents
  * [ ] Alternative scales (hexatonic, 19, 31, 50, ...)
* Chord
  * [ ] Parse from string
  * [ ] Modes
  * [ ] Comparisons
  * [ ] Generate from base note
  * [ ] Interval operations
  * [ ] Chord type conversions
  * [ ] Normalization
  * [ ] Relationships
  * [ ] Play on piano difficulty
  * [ ] Play on guitar difficulty
  * [ ] Relation: minorize, majorize, etc
  * [ ] Harmonies
  * [ ] More chords, more scales
* Progression
  * [ ] Parse from string
  * [ ] Unicode/Emoji output
  * [ ] Comparisons
  * [ ] Interval operations
  * [ ] Play on accordion difficulty
  * [ ] Play on Piano difficulty
  * [ ] Play on guitar difficulty
  * [ ] Shepard Tones
* Score
  * [ ] Unicode/Emoji output
  * [ ] SVG output
  * [ ] MIDI (un)marshaling
  * [ ] Interval operations
  * [ ] Comparisons
* Synth
  * [ ] Pythagorean Temperament
  * [ ] Equal Temperament
  * [ ] Werckmeister temperament
  * [ ] Mean-tone temperament
  * [ ] Generate sinusoidal
  * [ ] Parse sinusoidal
  * [ ] Sum sinusoidals
  * [ ] Hi-pass/low-pass filters
  * [ ] Synth filters
  * [ ] Augocorrelation
  * [ ] Phase shifting
  * [ ] Normalization
  * [ ] Buffer
  * [ ] Wav
  * [ ] Mp3
  * [ ] Write to sound card
  * [ ] OpenAL
  * [ ] Draw spectrum
  * [ ] Draw Waveform
  * [ ] Synth audio output
  * [ ] Frequency conversions
  * [ ] Frequency operations
  * [ ] Arpegiattors
* Other
  * [ ] Random generators
  * [ ] Tonnetz / Tonal Lattice support
  * [ ] Venn Diagram
  * [ ] Relationship SVG representations
  * [ ] Chime simulator
  * [ ] https://en.wikipedia.org/wiki/Musical_tuning
