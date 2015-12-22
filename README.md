Mzika
==========

A Go Video Music package.


## Integration Tests
As much as possible, a testing philosophy was followed in the development of this package. To that end, one can run the integration test which will verify that the core logic of the package is still working as expected:

```
$ go test

```

Note that the tests performed are real integration-level tests as they do make network requests. Consequently, it is possible to get spurious failures arising from network issues.