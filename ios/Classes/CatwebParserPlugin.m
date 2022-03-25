#import "CatwebParserPlugin.h"
#if __has_include(<catweb_parser/catweb_parser-Swift.h>)
#import <catweb_parser/catweb_parser-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "catweb_parser-Swift.h"
#endif

#import "libgo.h"


@implementation CatwebParserPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftCatwebParserPlugin registerWithRegistrar:registrar];
}


+ (void)fakeFunction {
    FakeCall();
}

@end
