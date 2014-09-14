
'use strict';

/* The angular application controllers */
var mycomputerControllers = angular.module('mycomputerControllers', []);

/* Todo: add comment */
mycomputerControllers.controller('HomeController', ['$scope',
  function($scope) {
}]);

/* This controller to get user and items from mycomputer beego api */
mycomputerControllers.controller('UserItemsController', ['$scope', '$routeParams', '$http',
  function($scope, $routeParams, $http) {
    /* Get the user data */
    $http.get('/api/' + $routeParams.username).success(function(data) {
      /* If the data is empty string, don't return objects */
      if(typeof data.Name == "undefined") {
        $scope.user = null;
      } else {
        $scope.user = data;
      }
    });

    /* Get the user items data */
    $http.get('/api/' + $routeParams.username +'/items').success(function(data) {
      /* If the data is empty string, don't return objects */
      if (typeof data[0].Username == "undefined") {
        $scope.items = null;
      } else {
        $scope.items = data;
      }
    });

    /* Determine whether pop up form to add item or not */
    $scope.adding = false;
  }
]);
