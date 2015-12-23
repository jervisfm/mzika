Mzika
==========

A Go Video Music package.


## Integration Tests
As much as possible, a testing philosophy was followed in the development of this package. To that end, one can run the integration test which will verify that the core logic of the package is still working as expected:

```
$ go test -v

```

Note that the tests performed are real integration-level tests as they do make network requests. Consequently, it is possible to get spurious failures arising from network issues.

## Dependencies
You need to have GopherJS installed inorder to generate the Javascript code.

## Browser Testing
You can manually test the generated Javascript function by opening the js/indx.html file in your Browser and using the JS Console.

```
# Build/Generate the Javascript
$ cd js
$ gopherjs build

# Start a Web Server
$ python -m SimpleHTTPServer

# Open page in browser
http://localhost:8000/index.html

# Test a query.
* Open JS Console
* Execute function below:
mzika.loadTopVideoJSONListingMostFavoritedToday(1, function(result, err) { console.log(result,err); });
# Should see output like below in Console
Object { Success: true, Result: Array[200] } null
```

## NodeJS Testing
There is also a JS file index.js that I had created to use for testing under NodeJS. However, as of the moment, this is not working because the net.Http requests fail to execute. I suspect this may be due to XHR somehow being unavaiable/inaccessible which is strange. Something to fix in the future as it will provide a nice avenue for quick feedback when making changes without having to rely on Browser. That said, the compiled Code Works fine in the target priemier platform which is a Browser JS execution engine which is what matters most.
