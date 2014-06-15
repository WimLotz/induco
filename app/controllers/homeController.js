(function () {
    var HomeController = function ($scope, $http, $window) {
        $scope.signIn = function () {

            $http.defaults.useXDomain = true;

            $http({method: 'POST', url: 'http://localhost:8080/googleAuthorise'})
                .success(function (data, status, headers, config) {
                    $window.location.href = data.url;
                })
                .error(function (data, status, headers, config) {
                    console.log('error');
                    console.log(status);
                    console.log(headers);
                    console.log(config);
                });
        };
    };
    HomeController.$inject = ['$scope', '$http', '$window'];
    angular.module('induco').controller('HomeController', HomeController);
}());