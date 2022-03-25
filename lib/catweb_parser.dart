import 'dart:typed_data';

import 'package:ffi/ffi.dart' as ffi;
import 'dart:ffi';
import 'dart:io';

import 'package:tuple/tuple.dart';

import 'gen/h/libgo.h.dart';

class NativeBinder {
  static final _native = NativeLibrary(Platform.isAndroid
      ? DynamicLibrary.open('libgo.so')
      : Platform.isWindows
          ? DynamicLibrary.open('./lib/libs/libgo.dll')
          : DynamicLibrary.process());

  static Uint8List parse(Uint8List raw) {
    final buffer = ffi.malloc.allocate<Int8>(raw.length + 1);
    buffer.asTypedList(raw.length).setAll(0, raw);

    final result = _native.ParseData(buffer, raw.length);
    final rsp = Uint8List.fromList(result.data.asTypedList(result.len));

    ffi.malloc.free(buffer);
    _native.FreeResult(result);
    return rsp;
  }

  static String runJs(Tuple2<String, String> input) {
    final jsBuffer = input.item1.toNativeUtf8().cast<Int8>();
    final inputBuffer = input.item2.toNativeUtf8().cast<Int8>();

    final result = _native.RunJs(jsBuffer, inputBuffer);

    final str = result.data.cast<ffi.Utf8>().toDartString(length: result.len);
    ffi.malloc.free(jsBuffer);
    ffi.malloc.free(inputBuffer);
    _native.FreeResult(result);
    return str;
  }
}
