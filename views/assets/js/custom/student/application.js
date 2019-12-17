$(document).ready(function(){

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

});