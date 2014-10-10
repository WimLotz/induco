(function () {

    var TagsController = function ($scope, tagsStorage) {
        $scope.tags = tagsStorage.getTags();

        $scope.addTag=function(item){
            tagsStorage.addTag(item);
        };

        $scope.closeTag = function (index) {
            tagsStorage.removeTag(index);
        };
    };

    var DashboardController = function ($scope) {
        $scope.message = "dashboard page";
    };

    var HomeController = function () {
    };

    var NavigationBarController = function ($scope, inducoApi) {
        $scope.signOut = function(){
            inducoApi.signOut();
        };
    };

    var SearchController = function ($scope) {
        $scope.message = 'search page';
    };

    var ProfileController = function ($scope, inducoApi, tagsStorage) {
        $scope.person = {};
        $scope.company = {};

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
            $scope.person.workExpTags = tagsStorage.getTags();

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

    ProfileController.$inject = ['$scope', 'inducoApi','tagsStorage'];
    DashboardController.$inject = ['$scope'];
    HomeController.$inject = [];
    NavigationBarController.$inject = ['$scope', 'inducoApi'];
    SearchController.$inject = ['$scope'];
    TagsController.$inject = ['$scope','tagsStorage'];

    angular.module("controllers", [])
        .controller('DashboardController', DashboardController)
        .controller('HomeController', HomeController)
        .controller('NavigationBarController', NavigationBarController)
        .controller('SearchController', SearchController)
        .controller('TagsController', TagsController)
        .controller('ProfileController', ProfileController);
})();