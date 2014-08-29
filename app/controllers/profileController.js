(function () {
    var ProfileController = function ($scope, $http) {
        var person = {};

        var createPerson = function () {
            person = {
                needHelp: $scope.person.needHelp,
                needWork: $scope.person.needWork,
                firstName: $scope.person.firstName,
                surname: $scope.person.surname,
                emailAddress: $scope.person.emailAddress
            }
            return person;
        };

        $scope.submitProfileForm = function (isValid) {
            if (isValid) {
                person = createPerson();
                savePersonProfile(person);
            }
        };

        var savePersonProfile = function (person) {
            $http({method: 'PUT', data: person, url: 'http://localhost:4567/saveProfile'})
                .success(function (data) {
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };
    };

    ProfileController.$inject = ['$scope', '$http'];

    angular.module('induco').controller('ProfileController', ProfileController);
}());