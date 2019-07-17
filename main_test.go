package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "mime/multipart"
    "os"
    "io"
)

func buildFileRequest(uri string, paramName string, path string) (*http.Request, error){
    file, _ := os.Open(path)
    defer file.Close()

    var body bytes.Buffer
    writer := multipart.NewWriter(&body)
    part, _ := writer.CreateFormFile(paramName,path)

    io.Copy(part, file)
    writer.Close()

    req, err:= http.NewRequest("POST", uri, &body)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    return req, err
}

func TestEchoHandler(t *testing.T) {

    req, _ := buildFileRequest("/echo", "file", "matrix.csv")
    rr:=httptest.NewRecorder()
    handler:= http.HandlerFunc(EchoHandler)
    handler.ServeHTTP(rr,req)

    var expected string = "1,2,3\n4,5,6\n7,8,9\n"
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    if rr.Body.String() != expected {
        t.Errorf("Wrong result, expected: %s,\n got %s", expected, rr.Body.String())
    }
}

func TestInvertHandler(t *testing.T) {

    req, _ := buildFileRequest("/invert", "file", "matrix.csv")
    rr:=httptest.NewRecorder()
    handler:= http.HandlerFunc(InvertHandler)
    handler.ServeHTTP(rr,req)

    var expected string = "1,4,7\n2,5,8\n3,6,9\n"
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    if rr.Body.String() != expected {
        t.Errorf("Wrong result, expected: %s,\n got %s", expected, rr.Body.String())
    }
}

func TestSumHandler(t *testing.T) {

    req, _ := buildFileRequest("/sum", "file", "matrix.csv")
    rr:=httptest.NewRecorder()
    handler:= http.HandlerFunc(SumHandler)
    handler.ServeHTTP(rr,req)

    var expected string = "45"
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    if rr.Body.String() != expected {
        t.Errorf("Wrong result, expected: %s,\n got %s", expected, rr.Body.String())
    }
}

func TestMultiplyHandler(t *testing.T) {

    req, _ := buildFileRequest("/multiply", "file", "matrix.csv")
    rr:=httptest.NewRecorder()
    handler:= http.HandlerFunc(MultiplyHandler)
    handler.ServeHTTP(rr,req)

    var expected string = "362880"
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    if rr.Body.String() != expected {
        t.Errorf("Wrong result, expected: %s,\n got %s", expected, rr.Body.String())
    }
}

func TestFlattenHandler(t *testing.T) {

    req, _ := buildFileRequest("/flatten", "file", "matrix.csv")
    rr:=httptest.NewRecorder()
    handler:= http.HandlerFunc(FlattenHandler)
    handler.ServeHTTP(rr,req)

    var expected string = "1,2,3,4,5,6,7,8,9"
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    if rr.Body.String() != expected {
        t.Errorf("Wrong result, expected: %s,\n got %s", expected, rr.Body.String())
    }
}