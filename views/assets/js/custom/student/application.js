$(document).ready(function(){

    $("#updateApplicantTypeBtn").hide();

    /** Matric form starts here **/

    $("form#matricInstitutionForm select#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $("form#matricInstitutionForm select#district");
        let townElement = $("form#matricInstitutionForm select#town");
        getDropDownElement(townElement, 'Town');
        let districtDropDown = getDropDownElement(districtElement, "District");
        populateLocationDropDown(districtDropDown, provinceId);
    });

    $("form#matricInstitutionForm select#district").change(function() {
        const districtId = $(this).val();
        let townElement = $("form#matricInstitutionForm select#town");
        let townDropDown = getDropDownElement(townElement, "Town");
        populateLocationDropDown(townDropDown, districtId);
    });

    $("form#matricInstitutionForm select#town").change(function() {
        $('form#matricInstitutionForm select#institutionType').prop('selectedIndex', 0);
        let institutionElement = $('form#matricInstitutionForm select#institution');
        getDropDownElement(institutionElement, "Institution");
    });

    $("form#matricInstitutionForm select#institutionType").change(function() {
        const institutionTypeId = $(this).val();
        const locationId = $('form#matricInstitutionForm select#town').val();
        let institutionElement = $("form#matricInstitutionForm select#institution");
        let institutionDropDown = getDropDownElement(institutionElement, "Institution");
        populateInstitutionDropDownByTypenLocation(institutionDropDown, institutionTypeId, locationId);
    });

    /** Matric form ends here **/

    /** Current institution form starts here **/

    $("#saveTertiaryCourseBtn").show();
    const selectedCourse = $("form#currentCourseForm select#course").val();
    if (selectedCourse) {
        $("#saveTertiaryCourseBtn").hide();
    }

    $("form#currentInstitutionForm select#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $("form#currentInstitutionForm select#district");
        let townElement = $("form#currentInstitutionForm select#town");
        getDropDownElement(townElement, 'Town');
        let districtDropDown = getDropDownElement(districtElement, "District");
        populateLocationDropDown(districtDropDown, provinceId);
    });

    $("form#currentInstitutionForm select#district").change(function() {
        const districtId = $(this).val();
        let townElement = $("form#currentInstitutionForm select#town");
        let townDropDown = getDropDownElement(townElement, "Town");
        populateLocationDropDown(townDropDown, districtId);
    });

    $("form#currentInstitutionForm select#town").change(function() {
        $('form#currentInstitutionForm select#institutionType').prop('selectedIndex', 0);
        let institutionElement = $('form#currentInstitutionForm select#institution');
        getDropDownElement(institutionElement, "Institution");
    });

    $("form#currentInstitutionForm select#institutionType").change(function() {
        const institutionTypeId = $(this).val();
        const locationId = $('form#currentInstitutionForm select#town').val();
        let institutionElement = $("form#currentInstitutionForm select#institution");
        let institutionDropDown = getDropDownElement(institutionElement, "Institution");
        populateInstitutionDropDownByTypenLocation(institutionDropDown, institutionTypeId, locationId);
    });

    /** Current institution form ends here **/

    /** Prospective institution form starts here **/

    $("#saveProspectiveTertiaryCourseBtn").show();
    const selectedProspectiveCourse = $("form#prospectiveCourseForm select#course").val();
    if (selectedProspectiveCourse) {
        $("#saveProspectiveTertiaryCourseBtn").hide();
    }

    $("form#prospectiveInstitutionForm select#province").change(function() {
        const provinceId = $(this).val();
        let districtElement = $("form#prospectiveInstitutionForm select#district");
        let townElement = $("form#prospectiveInstitutionForm select#town");
        getDropDownElement(townElement, 'Town');
        let districtDropDown = getDropDownElement(districtElement, "District");
        populateLocationDropDown(districtDropDown, provinceId);
    });

    $("form#prospectiveInstitutionForm select#district").change(function() {
        const districtId = $(this).val();
        let townElement = $("form#prospectiveInstitutionForm select#town");
        let townDropDown = getDropDownElement(townElement, "Town");
        populateLocationDropDown(townDropDown, districtId);
    });

    $("form#prospectiveInstitutionForm select#town").change(function() {
        $('form#prospectiveInstitutionForm select#institutionType').prop('selectedIndex', 0);
        let institutionElement = $('form#prospectiveInstitutionForm select#institution');
        getDropDownElement(institutionElement, "Institution");
    });

    $("form#prospectiveInstitutionForm select#institutionType").change(function() {
        const institutionTypeId = $(this).val();
        const locationId = $('form#prospectiveInstitutionForm select#town').val();
        let institutionElement = $("form#prospectiveInstitutionForm select#institution");
        let institutionDropDown = getDropDownElement(institutionElement, "Institution");
        populateInstitutionDropDownByTypenLocation(institutionDropDown, institutionTypeId, locationId);
    });

    /** Prospective institution form ends here **/

});

let checkSelectedApplicantType = function(initialApplicantType) {
    const selectedApplicantType = $("#applicantType").val();
    $("#updateApplicantTypeBtn").hide();
    if (initialApplicantType.localeCompare(selectedApplicantType)) {
        $("#updateApplicantTypeBtn").show();
    }
};

let checkSelectedCurrentCourse = function(initialSelectedCourse) {
    const selectedCourse = $("form#currentCourseForm select#course").val();
    $("#saveTertiaryCourseBtn").hide();
    if (initialSelectedCourse.localeCompare(selectedCourse)) {
        $("#saveTertiaryCourseBtn").show();
    }
};

let checkSelectedProspectiveCourse = function(initialSelectedCourse) {
    const selectedCourse = $("form#prospectiveCourseForm select#course").val();
    $("#saveProspectiveTertiaryCourseBtn").hide();
    if (initialSelectedCourse.localeCompare(selectedCourse)) {
        $("#saveProspectiveTertiaryCourseBtn").show();
    }
};