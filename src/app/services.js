(function () {
    var inducoApi = function ($http) {
        return {
            savePersonProfile: function (person) {
                return $http({method: 'POST', data: person, url: 'http://localhost:4567/savePersonProfile'});
            },
            saveCompanyProfile: function (company) {
                return $http({method: 'POST', data: company, url: 'http://localhost:4567/saveCompanyProfile'});
            },
            fetchCompanyProfiles:function(){
                return $http({method: 'GET', url: 'http://localhost:4567/fetchCompanyProfiles'});
            },
            fetchPersonProfiles:function(){
                return $http({method: 'GET', url: 'http://localhost:4567/fetchPersonProfiles'});
            },
            signOut:function(){
                return $http({method: 'POST', url: 'http://localhost:4567/googleDisconnect'});
            }
        }
    };

    inducoApi.$inject = ['$http'];

    angular.module("services", []).factory('inducoApi', inducoApi);;
})();