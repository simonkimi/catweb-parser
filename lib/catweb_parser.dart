
import 'dart:async';

import 'package:flutter/services.dart';

class CatwebParser {
  static const MethodChannel _channel = MethodChannel('catweb_parser');

  static Future<String?> get platformVersion async {
    final String? version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
}
