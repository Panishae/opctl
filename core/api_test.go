package core

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/opctl/engine/core/models"
)

var _ = Describe("_api", func() {
  Context(".AddSubOp() method", func() {
    It("should invoke compositionRoot.addSubOpUseCase.Execute() with expected args & return result", func() {

      /* arrange */
      providedAddSubOpReq := models.NewAddSubOpReq(&models.Url{}, "", "", "")

      // wire up fakes
      fakeAddSubOpUseCase := new(fakeAddSubOpUseCase)

      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.AddSubOpUseCaseReturns(fakeAddSubOpUseCase)

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.AddSubOp(*providedAddSubOpReq)

      /* assert */
      Expect(fakeAddSubOpUseCase.ExecuteArgsForCall(0)).To(Equal(*providedAddSubOpReq))
      Expect(fakeAddSubOpUseCase.ExecuteCallCount()).To(Equal(1))

    })
  })
  Context(".ListOps() method", func() {
    It("should invoke compositionRoot.listOpsUseCase.Execute() with expected args & return result", func() {

      /* arrange */
      providedProjectUrl := &models.Url{}
      expectedReturnedOps := make([]models.OpDetailedView, 0)

      // wire up fakes
      fakeListOpsUseCase := new(fakeListOpsUseCase)
      fakeListOpsUseCase.ExecuteReturns(expectedReturnedOps, nil)

      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.ListOpsUseCaseReturns(fakeListOpsUseCase)

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      actualReturnedOps, _ := objectUnderTest.ListOps(providedProjectUrl)

      /* assert */
      Expect(actualReturnedOps).To(Equal(expectedReturnedOps))

    })
  })
  Context(".RunOp() method", func() {
    It("should invoke compositionRoot.runOpUseCase.Execute() with expected args & return result", func() {

      /* arrange */
      providedRunOpReq := models.NewRunOpReq(&models.Url{}, map[string]string{})

      // wire up fakes
      fakeRunOpUseCase := new(fakeRunOpUseCase)

      fakeCompositionRoot := new(fakeCompositionRoot)
      fakeCompositionRoot.RunOpUseCaseReturns(fakeRunOpUseCase)

      objectUnderTest := &_api{
        compositionRoot:fakeCompositionRoot,
      }

      /* act */
      objectUnderTest.RunOp(*providedRunOpReq)

      /* assert */
      Expect(fakeRunOpUseCase.ExecuteArgsForCall(0)).To(Equal(*providedRunOpReq))
      Expect(fakeRunOpUseCase.ExecuteCallCount()).To(Equal(1))

    })
  })

})
