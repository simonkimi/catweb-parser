import 'package:flutter/services.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:catweb_parser/catweb_parser.dart';

void main() {
  const MethodChannel channel = MethodChannel('catweb_parser');

  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    channel.setMockMethodCallHandler((MethodCall methodCall) async {
      return '42';
    });
  });

  tearDown(() {
    channel.setMockMethodCallHandler(null);
  });

  test('getPlatformVersion', () async {
    expect(await CatwebParser.platformVersion, '42');
  });
}
