package pkg

import (
	"context"
	"github.com/golang-interfaces/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/pkg/manifest"
	"github.com/pkg/errors"
	"io/ioutil"
)

var _ = Context("pkg", func() {

	Context("GetManifest", func() {

		It("should call pkgHandle.GetContent w/ expected args", func() {
			/* arrange */
			providedPkgHandle := new(FakeHandle)
			// err to trigger immediate return
			providedPkgHandle.GetContentReturns(nil, errors.New("dummyError"))

			objectUnderTest := _Pkg{}

			/* act */
			objectUnderTest.GetManifest(providedPkgHandle)

			/* assert */
			actualCtx,
				actualPath := providedPkgHandle.GetContentArgsForCall(0)

			Expect(actualCtx).To(Equal(context.TODO()))
			Expect(actualPath).To(Equal(OpDotYmlFileName))
		})
		Context("pkgHandle.GetContent errs", func() {
			It("should return error", func() {
				/* arrange */
				getContentErr := errors.New("dummyError")

				providedPkgHandle := new(FakeHandle)
				// err to trigger immediate return
				providedPkgHandle.GetContentReturns(nil, getContentErr)

				objectUnderTest := _Pkg{}

				/* act */
				_, actualErr := objectUnderTest.GetManifest(providedPkgHandle)

				/* assert */
				Expect(actualErr).To(Equal(getContentErr))
			})
		})
		Context("pkgHandle.GetContent doesn't err", func() {
			It("should call ioUtil.ReadAll w/ expected args", func() {
				/* arrange */
				file, err := ioutil.TempFile("", "")
				if nil != err {
					panic(err)
				}

				providedPkgHandle := new(FakeHandle)
				providedPkgHandle.GetContentReturns(file, nil)

				fakeIOUtil := new(iioutil.Fake)
				// err to trigger immediate return
				fakeIOUtil.ReadAllReturns(nil, errors.New("dummyError"))

				objectUnderTest := _Pkg{
					ioUtil: fakeIOUtil,
				}

				/* act */
				objectUnderTest.GetManifest(providedPkgHandle)

				/* assert */
				actualReader := fakeIOUtil.ReadAllArgsForCall(0)

				Expect(actualReader).To(Equal(file))
			})
			Context("ioUtil.ReadAll errs", func() {
				It("should return error", func() {
					/* arrange */
					file, err := ioutil.TempFile("", "")
					if nil != err {
						panic(err)
					}

					providedPkgHandle := new(FakeHandle)
					providedPkgHandle.GetContentReturns(file, nil)

					readAllErr := errors.New("dummyError")
					fakeIOUtil := new(iioutil.Fake)
					fakeIOUtil.ReadAllReturns(nil, readAllErr)

					objectUnderTest := _Pkg{
						ioUtil: fakeIOUtil,
					}

					/* act */
					_, actualErr := objectUnderTest.GetManifest(providedPkgHandle)

					/* assert */
					Expect(actualErr).To(Equal(readAllErr))
				})
			})
			Context("ioUtil.ReadAll doesn't err", func() {
				It("should call manifest.Unmarshal w/ expected args & return result", func() {
					/* arrange */
					file, err := ioutil.TempFile("", "")
					if nil != err {
						panic(err)
					}

					providedPkgHandle := new(FakeHandle)
					providedPkgHandle.GetContentReturns(file, nil)

					bytesFromReadAll := []byte{2, 3}
					fakeIOUtil := new(iioutil.Fake)
					fakeIOUtil.ReadAllReturns(bytesFromReadAll, nil)

					expectedPkgManifest := &model.PkgManifest{
						Name: "dummyName",
					}
					expectedErr := errors.New("dummyError")
					fakeManifest := new(manifest.Fake)

					fakeManifest.UnmarshalReturns(expectedPkgManifest, expectedErr)

					objectUnderTest := _Pkg{
						ioUtil:   fakeIOUtil,
						manifest: fakeManifest,
					}

					/* act */
					actualPkgManifest, actualErr := objectUnderTest.GetManifest(providedPkgHandle)

					/* assert */
					Expect(fakeManifest.UnmarshalArgsForCall(0)).To(Equal(bytesFromReadAll))
					Expect(actualPkgManifest).To(Equal(expectedPkgManifest))
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
		})
	})
})