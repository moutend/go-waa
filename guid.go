package oleacc

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_CUIAutomation = ole.NewGUID("{ff48dba4-60ef-4201-aa87-54103eef594e}")

	IID_IAccessible                                  = ole.NewGUID("{618736e0-3c3d-11cf-810c-00aa00389b71}")
	IID_IUIAutomationSpreadsheetPattern              = ole.NewGUID("{7517a7c8-faae-4de9-9f08-29b91e8595c1}")
	IID_IUIAutomationElement                         = ole.NewGUID("{d22108aa-8ac5-49a5-837b-37bbb3d7591e}")
	IID_IUIAutomation                                = ole.NewGUID("{30cbe57d-d9d0-452a-ab13-7ac5ac4825ee}")
	IID_IUIAutomation2                               = ole.NewGUID("{34723aff-0c9d-49d0-9896-7ab52df8cd8a}")
	IID_IUIAutomation3                               = ole.NewGUID("{73D768DA-9B51-4B89-936E-C209290973E7}")
	IID_IUIAutomation4                               = ole.NewGUID("{1189C02A-05F8-4319-8E21-E817E3DB2860]")
	IID_IUIAutomationCacheRequest                    = ole.NewGUID("{b32a92b5-bc25-4078-9c08-d7ee95c48e03}")
	IID_IUIAutomationCondition                       = ole.NewGUID("{352ffba8-0973-437c-a61f-f64cafd81df9}")
	IID_IUIAutomationTreeWalker                      = ole.NewGUID("{4042c624-389c-4afc-a630-9df854a541fc}")
	IID_IUIAutomationEventHandler                    = ole.NewGUID("{146c3c17-f12e-4e22-8c27-f894b9b79c69}")
	IID_IUIAutomationPropertyChangedEventHandler     = ole.NewGUID("{40cd37d4-c756-4b0c-8c6f-bddfeeb13b50}")
	IID_IUIAutomationStructureChangedEventHandler    = ole.NewGUID("{e81d1b4e-11c5-42f8-9754-e7036c79f054}")
	IID_IUIAutomationFocusChangedEventHandler        = ole.NewGUID("{c270f6b5-5c69-4290-9745-7a7f97169468}")
	IID_IUIAutomationProxyFactory                    = ole.NewGUID("{85b94ecd-849d-42b6-b94d-d6db23fdf5a4}")
	IID_IUIAutomationProxyFactoryEntry               = ole.NewGUID("{d50e472e-b64b-490c-bca1-d30696f9f289}")
	IID_IUIAutomationProxyFactoryMapping             = ole.NewGUID("{09e31e18-872d-4873-93d1-1e541ec133fd}")
	IID_IUIAutomationTextEditTextChangedEventHandler = ole.NewGUID("{92FAA680-E704-4156-931A-E32D5BB38F3F}")
	IID_IUIAutomationChangesEventHandler             = ole.NewGUID("{58EDCA55-2C3E-4980-B1B9-56C17F27A2A0}")
	IID_IUIAutomationElement2                        = ole.NewGUID("{6749c683-f70d-4487-a698-5f79d55290d6}")
	IID_IUIAutomationElementArray                    = ole.NewGUID("{14314595-b4bc-4055-95f2-58f2e42c9855}")
	IID_IUIAutomationElement3                        = ole.NewGUID("{8471DF34-AEE0-4A01-A7DE-7DB9AF12C296}")
	IID_IUIAutomationElement4                        = ole.NewGUID("{3B6E233C-52FB-4063-A4C9-77C075C2A06B}")
	IID_IUIAutomationElement5                        = ole.NewGUID("{98141C1D-0D0E-4175-BBE2-6BFF455842A7}")
	IID_IUIAutomationElement6                        = ole.NewGUID("{4780d450-8bca-4977-afa5-a4a517f555e3}")
	IID_IUIAutomationElement7                        = ole.NewGUID("{204e8572-cfc3-4c11-b0c8-7da7420750b7}")
	IID_IUIAutomationTextRange                       = ole.NewGUID("{a543cc6a-f4ae-494b-8239-c814481187a8}")
	IID_IUIAutomationValuePattern                    = ole.NewGUID("{a94cd8b1-0844-4cd6-9d2d-640537ab39e9}")
	IID_IUIAutomationSelectionPattern                = ole.NewGUID("{5ed5202e-b2ac-47a6-b638-4b0bf140d78e}")
	IID_IUIAutomationAnnotationPattern               = ole.NewGUID("{9a175b21-339e-41b1-8e8b-623f6b681098}")
	IID_IUIAutomationScrollItemPattern               = ole.NewGUID("{b488300f-d015-4f19-9c29-bb595e3645ef}")
	IID_IUIAutomationTextChildPattern                = ole.NewGUID("{6552b038-ae05-40c8-abfd-aa08352aab86}")
	IID_IUIAutomationSpreadsheetItemPattern          = ole.NewGUID("{7d4fb86c-8d34-40e1-8e83-62c15204e335}")
	IID_IUIAutomationObjectModelPattern              = ole.NewGUID("{71c284b3-c14d-4d14-981e-19751b0d756d}")
	IID_IUIAutomationStylesPattern                   = ole.NewGUID("{85b5f0a2-bd79-484a-ad2b-388c9838d5fb}")
	IID_IUIAutomationTransformPattern                = ole.NewGUID("{a9b55844-a55d-4ef0-926d-569c16ff89bb}")
	IID_IUIAutomationTransformPattern2               = ole.NewGUID("{6d74d017-6ecb-4381-b38b-3c17a48ff1c2}")
	IID_IUIAutomationRangeValuePattern               = ole.NewGUID("{59213f4f-7346-49e5-b120-80555987a148}")
	IID_IUIAutomationTextRange2                      = ole.NewGUID("{BB9B40E0-5E04-46BD-9BE0-4B601B9AFAD4}")
	IID_IUIAutomationTextRange3                      = ole.NewGUID("{6A315D69-5512-4C2E-85F0-53FCE6DD4BC2}")
	IID_IUIAutomationTablePattern                    = ole.NewGUID("{620e691c-ea96-4710-a850-754b24ce2417}")
	IID_IUIAutomationVirtualizedItemPattern          = ole.NewGUID("{6ba3d7a6-04cf-4f11-8793-a8d1cde9969f}")
	IID_IUIAutomationLegacyIAccessiblePattern        = ole.NewGUID("{828055ad-355b-4435-86d5-3b51c14a9b1b}")
	IID_IUIAutomationSynchronizedInputPattern        = ole.NewGUID("{2233be0b-afb7-448b-9fda-3b378aa5eae1}")
	IID_IUIAutomationTableItemPattern                = ole.NewGUID("{0b964eb3-ef2e-4464-9c79-61d61737a27e}")
	IID_IUIAutomationDragPattern                     = ole.NewGUID("{1dc7b570-1f54-4bad-bcda-d36a722fb7bd}")
	IID_IUIAutomationWindowPattern                   = ole.NewGUID("{0faef453-9208-43ef-bbb2-3b485177864f}")
	IID_IUIAutomationBoolCondition                   = ole.NewGUID("{1b4e1f2e-75eb-4d0b-8952-5a69988e2307}")
	IID_                                             = ole.NewGUID("{a7d0af36-b912-45fe-9855-091ddc174aec}")
	IID_IUIAutomationSelectionItemPattern            = ole.NewGUID("{a8efa66a-0fda-421a-9194-38021f3578ea}")
	IID_IUIAutomationCustomNavigationPattern         = ole.NewGUID("{01EA217A-1766-47ED-A6CC-ACF492854B1F}")
	IID_IUIAutomationPropertyCondition               = ole.NewGUID("{99ebf2cb-5578-4267-9ad4-afd6ea77e94b}")
	IID_IUIAutomationGridPattern                     = ole.NewGUID("{414c3cdc-856b-4f5b-8538-3131c6302550}")
	IID_IUIAutomationInvokePattern                   = ole.NewGUID("{fb377fbe-8ea6-46d5-9c73-6499642d3059}")
)
