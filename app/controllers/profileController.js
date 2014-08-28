(function () {
    var ProfileController = function ($scope, $http) {
        var person = {};

        var createPerson = function () {
            person = {
                lookingForWork: $scope.person.lookingForWork,
                needWork: $scope.person.needWork,
                firstName: $scope.person.firstName,
                surname: $scope.person.surname,
                emailAddress: $scope.person.emailAddress
            }
           return person;
        };

        $scope.submitProfileForm = function (isValid) {
            if(isValid){
                person = createPerson();
                saveProfile(person);
            }
        };

        var saveProfile = function (person) {
            $http({method: 'PUT', data: person, url: 'http://localhost:4567/saveProfile'})
                .success(function (data) {
                    console.log(data);
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };
    };

    ProfileController.$inject = ['$scope', '$http'];

    angular.module('induco').controller('ProfileController', ProfileController);
}());