(function () {

    var ProfileController = function ($scope, $http) {
        $scope.person = {};
        $scope.company = {};
        $scope.person.workExpTags = [];

        $scope.addTag = function(tagName) {
            $scope.person.workExpTags.push({msg: tagName});
        };

        $scope.closeTag = function(index) {
            $scope.person.workExpTags.splice(index, 1);
        };

        $scope.submitPersonProfileForm = function (isValid) {
            if (isValid) {
                savePersonProfile();
            }
        };

        $scope.submitCompanyProfileForm = function (isValid) {
            if (isValid) {
                saveCompanyProfile();
            }
        };

        var saveCompanyProfile = function () {
            $http({method: 'PUT', data: $scope.company, url: 'http://localhost:4567/saveCompanyProfile'})
                .success(function (data) {
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        var savePersonProfile = function () {
            $http({method: 'PUT', data: $scope.person, url: 'http://localhost:4567/savePersonProfile'})
                .success(function (data) {
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        var fetchCompanyProfile = function () {
            $http({method: 'GET', url: 'http://localhost:4567/fetchCompanyProfile'})
                .success(function (data) {
                    $scope.company.name = data.name;
                    $scope.company.email = data.email;
                    $scope.company.telNumber = data.telNumber;
                    $scope.company.information = data.information;
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        var fetchPersonProfile = function () {
            $http({method: 'GET', url: 'http://localhost:4567/fetchPersonProfile'})
                .success(function (data) {
                    $scope.person.needHelp = data.needHelp;
                    $scope.person.needWork = data.needWork;
                    $scope.person.firstName = data.firstName;
                    $scope.person.surname = data.surname;
                    $scope.person.emailAddress = data.emailAddress;
                    $scope.person.personalInfo = data.personalInfo;
                    $scope.person.workExp = data.workExp;
                    $scope.person.requiredWork = data.requiredWork;
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        fetchPersonProfile();
//        fetchCompanyProfile();
    };

    ProfileController.$inject = ['$scope', '$http'];

    angular.module('induco').controller('ProfileController', ProfileController);
}());