(function () {
    var ProfileController = function ($scope, $http) {
        $scope.person = {};

        $scope.submitProfileForm = function (isValid) {
            if (isValid) {
                savePersonProfile();
            }
        };

        var savePersonProfile = function () {
            $http({method: 'PUT', data: $scope.person, url: 'http://localhost:4567/saveProfile'})
                .success(function (data) {
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        var fetchPersonProfile = function () {
            $http({method: 'GET', url: 'http://localhost:4567/fetchProfile'})
                .success(function (data) {
                    $scope.person.needHelp = data.needHelp;
                    $scope.person.needWork = data.needWork;
                    $scope.person.firstName = data.firstName;
                    $scope.person.surname = data.surname;
                    $scope.person.emailAddress = data.emailAddress;
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        fetchPersonProfile();
    };

    ProfileController.$inject = ['$scope', '$http'];

    angular.module('induco').controller('ProfileController', ProfileController);
}());