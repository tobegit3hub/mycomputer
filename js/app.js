'use strict';

/* App Module */

var phonecatApp = angular.module('phonecatApp', [
  'ngRoute',
  'phonecatControllers'
]);

phonecatApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/', {
        templateUrl: 'home.html',
        controller: 'HomeCtrl'
      }).
      when('/:phoneId', {
        templateUrl: 'tobe.html',
        controller: 'PhoneListCtrl'
      }).
      otherwise({
        redirectTo: '/'
      });
  }]);
