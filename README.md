go-papi
=======

AuthorityLabs Partner API Go Package

* Github: http://github.com/mtchavez/go_papi
* [![endorse](http://api.coderwall.com/mtchavez/endorsecount.png)](http://coderwall.com/mtchavez)

## Install

    go get github.com/mtchavez/go_papi

## Configure

Configure a client with your api key.

    client  := go_papi.Client { "asdf1234ApiKey" }

## Keywords

Keyword endpoint calls.

### POST

Post a keyword which defaults to Google and en-US locale. *TODO: Allow all parameters*

    client  := go_papi.Client { "6rvrqCVu6GuPKZVQUFAF" }
    client.KeywordsPost("my keyword")

### Priority POST

Same as POST but added to a priority queue. *TODO: Allow all parameters*

    client  := go_papi.Client { "6rvrqCVu6GuPKZVQUFAF" }
    client.KeywordsPriorityPost("my keyword")

### GET

Get your SERP results for a keyword with default params. *TODO: Allow all parameters*
Returns a simplejson.Json interface.

    client  := go_papi.Client { "6rvrqCVu6GuPKZVQUFAF" }
    json2, _ := client.KeywordsGet("my keyword")
    
    // Get first ranking in serp
    fmt.Println(json2.Get("serp").Get("1"))

## TODO

* Allow all valid parameters to be passed to API
* Implement missing endpoints
* Add tests

## License

Written by Chavez

Released under the MIT License: http://www.opensource.org/licenses/mit-license.php
