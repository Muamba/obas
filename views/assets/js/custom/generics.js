const LOCATION_RESTAPI = "http://localhost:4000/location/api/";
const INSTITUTION_RESTAPI = "http://localhost:4000/institution/api/";

let populateLocationDropDown = function(element, locationId) {
    if (locationId) {
        const url = LOCATION_RESTAPI + "getforparent/" + locationId;
        $.get(url, function(data) {
            $.each(data, function (key, value) {
                let option = new Option(value.name, value.locationId);
                element.append(option);
            });
        });
    }
};

let getDropDownElement = function(element, displayText) {
    let option = '<option value="" disabled selected>Select '+ displayText +'</option>';
    element.empty();
    element.append(option);
    return element;
};

let populateInstitutionDropDownByLocation = function (element, locationId) {
    if (locationId) {
        const url = INSTITUTION_RESTAPI + "getInstitutionsInLocation/" + locationId;
        $.get(url, function(institutions) {
            $.each(institutions, function (key, value) {
                let option = new Option(value.name, value.id);
                element.append(option);
            });
        });
    }
};

let populateInstitutionDropDownByType = function (element, institutionTypeId) {
    if (institutionTypeId) {
        const url = INSTITUTION_RESTAPI + "getInstitutionsByType/" + institutionTypeId;
        $.get(url, function(institutions) {
            $.each(institutions, function (key, value) {
                let option = new Option(value.name, value.id);
                element.append(option);
            });
        });
    }
};

let populateInstitutionDropDownByTypenLocation = function (element, institutionTypeId, locationId) {
    if (locationId && institutionTypeId) {
        const url = INSTITUTION_RESTAPI + "getInstitutionsByTypenLocation/" + institutionTypeId + "/" + locationId;
        $.get(url, function(institutions) {
            $.each(institutions, function (key, value) {
                let option = new Option(value.name, value.id);
                element.append(option);
            });
        });
    }
};