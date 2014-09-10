(function () {

    var ProfileController = function ($scope, inducoApi) {
        $scope.person = {};
        $scope.company = {};
        $scope.person.workExpTags = [];

        $scope.addTag = function (tagName) {
            $scope.person.workExpTags.push({msg: tagName});
        };

        $scope.closeTag = function (index) {
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
            inducoApi.saveCompanyProfile($scope.company)
                .error(function (data) {
                    console.log("Error occurred trying to save a company: " + data);
                });
        };

        var savePersonProfile = function () {
            inducoApi.savePersonProfile($scope.person)
                .error(function (data) {
                    console.log("Error occurred trying to save a person: " + data);
                });
        };

        var fetchCompanyProfiles = function () {
            inducoApi.fetchCompanyProfiles()
                .success(function (data) {
                    if (data.length > 0) {
                        $scope.company.id = data[0].id
                        $scope.company.name = data[0].name;
                        $scope.company.email = data[0].email;
                        $scope.company.telNumber = data[0].telNumber;
                        $scope.company.information = data[0].information;
                    }
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        var fetchPersonProfiles = function () {
            inducoApi.fetchPersonProfiles()
                .success(function (data) {
                    if (data.length > 0) {
                        $scope.person.id = data[0].id;
                        $scope.person.needHelp = data[0].needHelp;
                        $scope.person.needWork = data[0].needWork;
                        $scope.person.firstName = data[0].firstName;
                        $scope.person.surname = data[0].surname;
                        $scope.person.emailAddress = data[0].emailAddress;
                        $scope.person.personalInfo = data[0].personalInfo;
                        $scope.person.workExp = data[0].workExp;
                        $scope.person.requiredWork = data[0].requiredWork;
                    }
                })
                .error(function (data) {
                    console.log('An error has occurred: ' + data);
                });
        };

        fetchPersonProfiles();
        fetchCompanyProfiles();
    };

    ProfileController.$inject = ['$scope', 'inducoApi'];

    angular.module('induco').controller('ProfileController', ProfileController);
}());