# mockvar

[![Build Status](https://travis-ci.org/nowk/mockvar.svg?branch=master)][0]
[![GoDoc](https://godoc.org/gopkg.in/nowk/mockvar.v0?status.svg)][1]

  [0]: https://travis-ci.org/nowk/mockvar
  [1]: http://godoc.org/gopkg.in/nowk/mockvar.v0

Mock and restore var assignments

## Install

    go get gopkg.in/nowk/mockvar.v0


## Usage

Basic use is to redefine var assigned values then restore them to their original values when requiring a simple way to mock out certain code for tests.

Given:

    var getProfile = func() (*Profile, error) {
      return &Profile{
        Name: "Batman",
      }, nil
    }

    func SaveProfile() error {
      p, err := getProfile()
      if err != nil {
        return err
      }

      return nil
    }

To mock a returned error:

    restore := mockvar.Mock(&getProfile, func() (*Profile, error) {
      return nil, errors.New("Oops!")
    })
    defer restore() // restore back to original

    err := SaveProfile()
    assert.Equal(t, "Oops!", err.Error())


### License

MIT
