package rest

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http/httptest"
  "net/http"
  "strings"
  "github.com/chrisdostert/mux"
)

var _ = Describe("runPipelineHandler", func() {

  Context("ServeHTTP() method", func() {
    It("should return StatusCode of 400 if projectUrl is malformed in Request", func() {

      /* arrange */
      objectUnderTest := runPipelineHandler{}
      recorder := httptest.NewRecorder()
      m := mux.NewRouter()
      m.Handle(runPipelineRelUrlTemplate, objectUnderTest)

      providedProjectUrl := "%%invalidProjectUrl%%"
      httpReq, err := http.NewRequest(http.MethodGet, "", nil)
      if (nil != err) {
        Fail(err.Error())
      }

      // brute force a request with malformed projectUrl
      httpReq.URL.Path = strings.Replace(runPipelineRelUrlTemplate, "{projectUrl}", providedProjectUrl, 1)

      /* act */
      m.ServeHTTP(recorder,httpReq)

      /* assert */
      Expect(recorder.Code).To(Equal(http.StatusBadRequest))

    })
  })
})