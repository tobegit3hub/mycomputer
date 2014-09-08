function PhoneListCtrl($scope) {
  $scope.phones = [
    {"name": "Nexus S",
     "image_url": "macbook_air.jpeg",
     "snippet": "Fast just got faster with Nexus S."},
    {"name": "Motorola XOOM™ with Wi-Fi",
     "image_url": "miwifi.jpg",
     "snippet": "The Next, Next Generation tablet."},
    {"name": "MOTOROLA XOOM™",
     "image_url": "newbalance.jpg",
     "snippet": "The Next, Next Generation tablet."}
  ];
}

/*
var app = angular.module("app", []);

app.controller("PhoneListCtrl", function($scope, $http) {
  $http.get('https://raw.githubusercontent.com/tobegit3hub/mycomputer/master/phones.json').
    success(function(data, status, headers, config) {
      $scope.posts = data;
    }).
    error(function(data, status, headers, config) {
      // log error
    });
});
*/



