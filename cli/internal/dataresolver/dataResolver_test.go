package dataresolver

import (
	"errors"
	"fmt"

	"github.com/golang-interfaces/ios"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/cli/internal/cliexiter"
	"github.com/opctl/opctl/cli/internal/cliparamsatisfier"
	cliModel "github.com/opctl/opctl/cli/internal/model"
	"github.com/opctl/opctl/cli/internal/nodeprovider"
	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/api/client"
)

var _ = Context("dataResolver", func() {
	Context("Resolve", func() {
		Context("os.Getwd errs", func() {
			It("should call exiter w/ expected args", func() {
				/* arrange */
				fakeIOS := new(ios.Fake)
				expectedError := errors.New("dummyError")
				fakeIOS.GetwdReturns("", expectedError)

				fakeCliExiter := new(cliexiter.Fake)

				objectUnderTest := _dataResolver{
					cliExiter: fakeCliExiter,
					os:        fakeIOS,
				}

				/* act */
				objectUnderTest.Resolve(
					"dummyDataRef",
					nil,
				)

				/* assert */
				Expect(fakeCliExiter.ExitArgsForCall(0)).
					To(Equal(cliexiter.ExitReq{Message: expectedError.Error(), Code: 1}))
			})
		})
		Context("os.Getwd doesn't err", func() {
			It("should call data.NewFSProvider w/ expected args", func() {
				/* arrange */
				fakeAPIClient := new(client.Fake)
				fakeNodeHandle := new(cliModel.FakeNodeHandle)
				fakeNodeHandle.APIClientReturns(fakeAPIClient)

				fakeNodeProvider := new(nodeprovider.Fake)
				fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, errors.New("dummyError"))

				fakeData := new(data.Fake)
				fakeFSDataProvider := new(data.FakeProvider)
				fakeData.NewFSProviderReturns(fakeFSDataProvider)

				fakeIOS := new(ios.Fake)
				workDir := "dummyWorkDir"
				fakeIOS.GetwdReturns(workDir, nil)

				objectUnderTest := _dataResolver{
					data:         fakeData,
					cliExiter:    new(cliexiter.Fake),
					os:           fakeIOS,
					nodeProvider: fakeNodeProvider,
				}

				/* act */
				objectUnderTest.Resolve("dummyDataRef", &model.PullCreds{})

				/* assert */
				Expect(fakeData.NewFSProviderArgsForCall(0)).To(ConsistOf(workDir))
			})
			It("should call data.NewNodeProvider w/ expected args", func() {
				/* arrange */
				providedPullCreds := &model.PullCreds{
					Username: "dummyUsername",
					Password: "dummyPassword",
				}

				fakeAPIClient := new(client.Fake)
				fakeNodeHandle := new(cliModel.FakeNodeHandle)
				fakeNodeHandle.APIClientReturns(fakeAPIClient)

				fakeNodeProvider := new(nodeprovider.Fake)
				fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

				fakeData := new(data.Fake)
				fakeNodeDataProvider := new(data.FakeProvider)
				fakeData.NewNodeProviderReturns(fakeNodeDataProvider)

				// err to trigger immediate return
				fakeData.ResolveReturns(nil, errors.New("dummyError"))

				objectUnderTest := _dataResolver{
					data:         fakeData,
					cliExiter:    new(cliexiter.Fake),
					os:           new(ios.Fake),
					nodeProvider: fakeNodeProvider,
				}

				/* act */
				objectUnderTest.Resolve(
					"dummyDataRef",
					providedPullCreds,
				)

				/* assert */
				actualAPIClient,
					actualPullCreds := fakeData.NewNodeProviderArgsForCall(0)
				Expect(actualAPIClient).To(Equal(fakeAPIClient))
				Expect(actualPullCreds).To(Equal(providedPullCreds))
			})
			It("should call data.Resolve w/ expected args", func() {
				/* arrange */
				providedDataRef := "dummyDataRef"

				fakeAPIClient := new(client.Fake)
				fakeNodeHandle := new(cliModel.FakeNodeHandle)
				fakeNodeHandle.APIClientReturns(fakeAPIClient)

				fakeNodeProvider := new(nodeprovider.Fake)
				fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

				fakeData := new(data.Fake)

				fakeFSDataProvider := new(data.FakeProvider)
				fakeData.NewFSProviderReturns(fakeFSDataProvider)

				fakeNodeDataProvider := new(data.FakeProvider)
				fakeData.NewNodeProviderReturns(fakeNodeDataProvider)

				// err to trigger immediate return
				fakeData.ResolveReturns(nil, errors.New("dummyError"))

				objectUnderTest := _dataResolver{
					data:         fakeData,
					cliExiter:    new(cliexiter.Fake),
					os:           new(ios.Fake),
					nodeProvider: fakeNodeProvider,
				}

				/* act */
				objectUnderTest.Resolve(
					providedDataRef,
					&model.PullCreds{},
				)

				/* assert */
				_,
					actualDataRef,
					actualProviders := fakeData.ResolveArgsForCall(0)

				Expect(actualDataRef).To(Equal(providedDataRef))
				Expect(actualProviders).To(ConsistOf(fakeFSDataProvider, fakeNodeDataProvider))
			})
			Context("data.Resolve errs", func() {
				Context("data.ErrDataProviderAuthorization", func() {
					It("should call cliParamSatisfier.Satisfy w/ expected args", func() {
						/* arrange */
						fakeAPIClient := new(client.Fake)
						fakeNodeHandle := new(cliModel.FakeNodeHandle)
						fakeNodeHandle.APIClientReturns(fakeAPIClient)

						fakeNodeProvider := new(nodeprovider.Fake)
						fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

						fakeData := new(data.Fake)
						expectedError := model.ErrDataProviderAuthorization{}
						fakeData.ResolveReturnsOnCall(0, nil, expectedError)

						username := "dummyUsername"
						password := "dummyPassword"

						fakeCliParamSatisfier := new(cliparamsatisfier.Fake)
						fakeCliParamSatisfier.SatisfyReturns(
							map[string]*model.Value{
								usernameInputName: {String: &username},
								passwordInputName: {String: &password},
							},
						)

						objectUnderTest := _dataResolver{
							data:              fakeData,
							cliParamSatisfier: fakeCliParamSatisfier,
							cliExiter:         new(cliexiter.Fake),
							os:                new(ios.Fake),
							nodeProvider:      fakeNodeProvider,
						}

						/* act */
						objectUnderTest.Resolve("dummyDataRef", &model.PullCreds{})

						/* assert */
						_, actualInputs := fakeCliParamSatisfier.SatisfyArgsForCall(0)
						Expect(actualInputs).To(Equal(credsPromptInputs))
					})
					It("should call data.NewNodeProvider w/ expected args", func() {
						/* arrange */
						fakeAPIClient := new(client.Fake)
						fakeNodeHandle := new(cliModel.FakeNodeHandle)
						fakeNodeHandle.APIClientReturns(fakeAPIClient)

						fakeNodeProvider := new(nodeprovider.Fake)
						fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

						fakeData := new(data.Fake)
						fakeNodeDataProvider := new(data.FakeProvider)
						fakeData.NewNodeProviderReturns(fakeNodeDataProvider)

						expectedError := model.ErrDataProviderAuthentication{}
						fakeData.ResolveReturnsOnCall(0, nil, expectedError)

						pullCreds := &model.PullCreds{
							Username: "dummyUsername",
							Password: "dummyPassword",
						}

						fakeCliParamSatisfier := new(cliparamsatisfier.Fake)
						fakeCliParamSatisfier.SatisfyReturns(
							map[string]*model.Value{
								usernameInputName: {String: &pullCreds.Username},
								passwordInputName: {String: &pullCreds.Password},
							},
						)

						objectUnderTest := _dataResolver{
							data:              fakeData,
							cliParamSatisfier: fakeCliParamSatisfier,
							cliExiter:         new(cliexiter.Fake),
							os:                new(ios.Fake),
							nodeProvider:      fakeNodeProvider,
						}

						/* act */
						objectUnderTest.Resolve(
							"dummyDataRef",
							&model.PullCreds{},
						)

						/* assert */
						actualAPIClient,
							actualPullCreds := fakeData.NewNodeProviderArgsForCall(1)
						Expect(actualAPIClient).To(Equal(fakeAPIClient))
						Expect(actualPullCreds).To(Equal(pullCreds))
					})
					It("should call data.Resolve w/ expected args", func() {
						/* arrange */
						providedDataRef := "dummyDataRef"

						fakeAPIClient := new(client.Fake)
						fakeNodeHandle := new(cliModel.FakeNodeHandle)
						fakeNodeHandle.APIClientReturns(fakeAPIClient)

						fakeNodeProvider := new(nodeprovider.Fake)
						fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

						fakeData := new(data.Fake)

						fakeFSDataProvider := new(data.FakeProvider)
						fakeData.NewFSProviderReturns(fakeFSDataProvider)

						fakeNodeDataProvider := new(data.FakeProvider)
						fakeData.NewNodeProviderReturns(fakeNodeDataProvider)

						expectedError := model.ErrDataProviderAuthentication{}
						fakeData.ResolveReturnsOnCall(0, nil, expectedError)

						fakeCliParamSatisfier := new(cliparamsatisfier.Fake)
						fakeCliParamSatisfier.SatisfyReturns(
							map[string]*model.Value{
								usernameInputName: {String: new(string)},
								passwordInputName: {String: new(string)},
							},
						)

						objectUnderTest := _dataResolver{
							data:              fakeData,
							cliParamSatisfier: fakeCliParamSatisfier,
							cliExiter:         new(cliexiter.Fake),
							os:                new(ios.Fake),
							nodeProvider:      fakeNodeProvider,
						}

						/* act */
						objectUnderTest.Resolve(
							providedDataRef,
							&model.PullCreds{},
						)

						/* assert */
						_,
							actualDataRef,
							actualProviders := fakeData.ResolveArgsForCall(1)

						Expect(actualDataRef).To(Equal(providedDataRef))
						Expect(actualProviders).To(ConsistOf(fakeFSDataProvider, fakeNodeDataProvider))
					})
				})
				Context("not data.ErrAuthenticationFailed", func() {
					It("should call exiter w/ expected args", func() {
						/* arrange */
						providedDataRef := "dummyDataRef"

						fakeAPIClient := new(client.Fake)
						fakeNodeHandle := new(cliModel.FakeNodeHandle)
						fakeNodeHandle.APIClientReturns(fakeAPIClient)

						fakeNodeProvider := new(nodeprovider.Fake)
						fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

						fakeData := new(data.Fake)
						resolveError := errors.New("dummyError")
						fakeData.ResolveReturns(nil, resolveError)

						expectedError := fmt.Sprintf(
							"Unable to resolve pkg '%v'; error was %v",
							providedDataRef,
							resolveError.Error(),
						)

						fakeCliExiter := new(cliexiter.Fake)

						objectUnderTest := _dataResolver{
							data:         fakeData,
							cliExiter:    fakeCliExiter,
							os:           new(ios.Fake),
							nodeProvider: fakeNodeProvider,
						}

						/* act */
						objectUnderTest.Resolve("dummyDataRef", &model.PullCreds{})

						/* assert */
						Expect(fakeCliExiter.ExitArgsForCall(0)).
							To(Equal(cliexiter.ExitReq{Message: expectedError, Code: 1}))

					})
				})
			})
			Context("data.Resolve doesn't err", func() {
				It("should return expected result", func() {
					/* arrange */
					fakeAPIClient := new(client.Fake)
					fakeNodeHandle := new(cliModel.FakeNodeHandle)
					fakeNodeHandle.APIClientReturns(fakeAPIClient)

					fakeNodeProvider := new(nodeprovider.Fake)
					fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)

					fakeData := new(data.Fake)
					fakeHandle := new(data.FakeHandle)

					// err to trigger immediate return
					fakeData.ResolveReturns(fakeHandle, nil)

					objectUnderTest := _dataResolver{
						data:         fakeData,
						os:           new(ios.Fake),
						nodeProvider: fakeNodeProvider,
					}

					/* act */
					actualPkgHandle := objectUnderTest.Resolve(
						"dummyDataRef",
						&model.PullCreds{},
					)

					/* assert */
					Expect(actualPkgHandle).To(Equal(fakeHandle))
				})
			})
		})
	})
})
