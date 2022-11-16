const url = new URL(window.location.href);
const campaignUrlName = url.href.split(':8080/')[1];

// getting the inputs 
const form = document.getElementsByClassName("campaign__form");
const campaignName = document.getElementById("campaignName");
const campaignStart = document.getElementById("campaignStart");
const campaignEnd = document.getElementById("campaignEnd");
const campaignPrice = document.getElementById("campaignPrice");
const campaignObjective = document.getElementById("campaignObjective");
const campaignPricePerDisplay = document.getElementById("campaignPricePerDisplay");
const submit = document.getElementById("submit");
const deleteButton = document.getElementById("deleteButton");


fetch("http://localhost:8080/api/campaign/getByName/" + campaignUrlName)
    .then((res) => {return res.json()})
    .then((data) => {
        campaignName.setAttribute("value", data.name);
        campaignStart.setAttribute("value", data.startDate);
        campaignEnd.setAttribute("value", data.endDate);
        campaignPrice.setAttribute("value", data.price);
        campaignObjective.setAttribute("value", data.objective);
        campaignPricePerDisplay.setAttribute("value", data.pricePerDisplay);
    });

campaignObjective.addEventListener("change", () => {
    if ((campaignPrice.value && campaignObjective.value) !== undefined) {
        campaignPricePerDisplay.value = parseInt(campaignPrice.value) / parseInt(campaignObjective.value);
    }
});

// function to request to update
const updateCampaign = () => {
    const nameValue = campaignName.value;
    const startValue = campaignStart.value;
    const endValue = campaignEnd.value;
    const priceValue = parseInt(campaignPrice.value);
    const objectiveValue = parseInt(campaignObjective.value);
    const pricePerDisplay = parseInt(campaignPricePerDisplay.value);

    const updatedInfos = {
        "name": nameValue,
        "startDate": startValue,
        "endDate": endValue,
        "price": priceValue,
        "objective": objectiveValue,
        "pricePerDisplay": pricePerDisplay
    }

    fetch("http://localhost:8080/api/campaign/update/" + campaignUrlName, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedInfos)
    })
        .then((res) => {
            return res.json();
        })
        .then((data) => {
            console.log(data);
        });
}

// update request send when click on submit button
submit.addEventListener("click", (event) => {
    event.preventDefault();
    updateCampaign();
})

// function to delete the campaign 
const deleteCampaign = () => {
    fetch("http://localhost:8080/api/campaign/delete/" + campaignUrlName, {
        method: "DELETE"
    })
        .then((res) => {
            return res.json();
        })
        .then((data) => {
            console.log(data);
        })
}

deleteButton.addEventListener("click" , () => {
    deleteCampaign();
});