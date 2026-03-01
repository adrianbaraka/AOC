#!/usr/bin/env dart

import '../../../utils/utils.dart';

/// Return the index of the max element in the list
int getidxmax(List<int> l) {
  int max = l[0];
  int idx = 0;

  for (var i = 0; i < l.length; i++) {
    if (l[i] > max) {
      max = l[i];
      idx = i;
    }
  }
  return idx;
}

void main(List<String> args) {
  // get and convert to int
  var data = loadData(args)[0];
  var pattern = RegExp(r"\s+");
  var sdata = data.trim().split(pattern);

  List<int> dataInt = [];

  for (int i = 0; i < sdata.length; i++) {
    int num = toInt(sdata[i]);
    dataInt.add(num);
  }

  Map<String, int> seen = {};

  seen[dataInt.join("")] = 0;

  int cycles = 1;
  while (true) {
    // get the index of the max
    var idx = getidxmax(dataInt);
    int j = idx + 1;
    var val = dataInt[idx];

    // remove all from idx
    dataInt[idx] = 0;

    // loop distributing val
    while (val > 0) {
      if (j >= dataInt.length) {
        j = 0;
      }
      // add to j remove from val
      dataInt[j]++;
      val--;
      j++;
    }

    // check if already seen
    var next = dataInt.join("");
    final lastseen = seen[next];
    if (lastseen != null) {
      print("Part 2: ${cycles - lastseen}");
      break;
    }

    seen[next] = cycles;
    cycles++;
    //print(seen);
  }
  print("Part 1: $cycles");
}
