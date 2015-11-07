# Status of this fork

Development by the original Fresh creator seems to have slowed down a lot with important pull requests waiting for many months to be reviewed and merged.

I will be cherry-picking commits from all the forks just to have a better, more up to date version. Unless I stumble upon something affecting me personally I don't intend to put significant amount of time into improving this already great tool.

I promise to be very responsive reviewing and accepting (or rejecting) pull requests.

# Fresh

Fresh is a command line tool that builds and (re)starts your web application every time you save a Go or template file.

If the web framework you are using supports the Fresh runner, it will show build errors on your browser.

It has been tested with:
* [chi](https://github.com/pressly/chi)
* [goji](https://github.com/zenazn/goji)
* [gocraft/web](https://github.com/gocraft/web)
* [Martini](https://github.com/codegangsta/martini)
* [Traffic](https://github.com/pilu/traffic)

## Installation

    go get github.com/c2h5oh/fresh

## Usage

    cd /path/to/myapp

Start fresh:

    fresh

Fresh will watch for file events, and every time you create/modifiy/delete a file it will build and restart the application.
If `go build` returns an error, it will log it in the tmp folder.

[Traffic](https://github.com/pilu/traffic) already has a middleware that shows the content of that file if it is present. This middleware is automatically added if you run a Traffic web app in dev mode with Fresh.
Check the `_examples` folder if you want to use it with Martini or Gocraft Web.

You can use the `-c` options if you want to specify a config file:

    fresh -c runner.conf

Here is a sample config file with the default settings:

    root:              .
    tmp_path:          ./tmp
    build_name:        runner-build
    build_log:         runner-build-errors.log
    valid_ext:         .go, .tpl, .tmpl, .html
    build_delay:       600
    colors:            1
    log_color_main:    cyan
    log_color_build:   yellow
    log_color_runner:  green
    log_color_watcher: magenta
    log_color_app:

## Original Author

* [Andrea Franz](http://gravityblast.com)

## Maintainter of this fork

* [Maciej Lisiewski](https://twitter.com/lisiewski)


## More

* [Mailing List](https://groups.google.com/d/forum/golang-fresh)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
