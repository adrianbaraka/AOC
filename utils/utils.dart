import 'dart:io';

/// Read from stdin or the passed path and returns a list of the contents.
List<String> loadData(List<String> args) {
  try {
    // read from 1st cl arguement
    if (args.length > 0 && args[0] != "-") {
      File file = File(args[0]);
      var data = file.readAsLinesSync();
      //print(data);
      return data;
    }

    // read from stdin
    List<String> lines = [];
    while (true) {
      var line = stdin.readLineSync();
      if (line == null) {
        break;
      }
      lines.add(line);
    }
    return lines;
  } on Exception catch (e) {
    stderr.writeln("An error occurred. \n$e");
    exit(1);
  }
}

/// Converts the string str to int. If it can't fatal
int toInt(String str) {
    var num = int.tryParse(str);
    if (num == null) {
        stderr.writeln("Error: Cannot convert '$str' to an int");
        exit(1);
    }
    return num;
}