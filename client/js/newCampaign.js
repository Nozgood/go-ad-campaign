// getting the inputs 
const form = document.getElementsByClassName("campaign__form");
const campaignName = document.getElementById("campaignName");
const campaignStart = document.getElementById("campaignStart");
const campaignEnd = document.getElementById("campaignEnd");
const campaignPrice = document.getElementById("campaignPrice");
const campaignObjective = document.getElementById("campaignObjective");
const campaignPricePerDisplay = document.getElementById("campaignPricePerDisplay");
const submit = document.getElementById("submit");

// fetch to post the new campaign
const newCampaign = () => {
    const nameValue = campaignName.value;
    const startValue = campaignStart.value;
    const endValue = campaignEnd.value;
    const priceValue = parseInt(campaignPrice.value);
    const objectiveValue = parseInt(campaignObjective.value);
    const pricePerDisplay = parseInt(priceValue/objectiveValue);

    const formSubmit = {
        "name": nameValue,
        "startDate": startValue,
        "endDate": endValue,
        "price": priceValue,
        "objective": objectiveValue,
        "pricePerDisplay": pricePerDisplay
    }

    fetch("http://localhost:8080/api/campaign/create", {
        method: "POST",
        headers: {
            "Content-Type": 'application/json',
        },
        body: JSON.stringify(formSubmit),
    })
    .then((res) => {
        return res.json();
    })
    .then((data) => {
        console.log(data);
    })
}


// display the pricePerDisplay
campaignObjective.addEventListener("change", () => {
    if ((campaignPrice.value && campaignObjective.value) !== undefined) {
        campaignPricePerDisplay.value = parseInt(campaignPrice.value) / parseInt(campaignObjective.value);
    }
})

// event when submit 
submit.addEventListener("click", (event) => {
    event.preventDefault();
    newCampaign();
})
