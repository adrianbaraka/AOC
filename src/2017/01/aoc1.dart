#!/usr/bin/env dart
import "dart:io";
import '../../../utils/utils.dart';

void main(List<String> args) {
  var data = loadData(args)[0];
  part1(data);
  part2(data);
}

// function to add to the sum
int addSum(String str, int sum) {
  var num = int.tryParse(str);
  if (num == null) {
    //fatal
    stderr.writeln("Cannot convert $str to int.");
    exit(1);
  }
  sum += num;

  return sum;
}

void part1(String data) {
  var sum = 0;
  for (var i = 0; i < data.length - 1; i++) {
    // part 1
    if (data[i] == data[i + 1]) {
      sum = addSum(data[i], sum);
    }
    if (i == data.length - 2 && data[i + 1] == data[0]) {
      sum = addSum(data[i + 1], sum);
    }
  }

  print("Part 1: $sum");
}

void part2(String data) {
  var sum = 0;
  int half = (data.length / 2).toInt();

  for (int i = 0; i < data.length; i++) {
    int next = (i + half) % data.length;
    if (data[i] == data[next]) {
      sum = addSum(data[i], sum);
    }
  }

  print("Part 2: $sum");
}
