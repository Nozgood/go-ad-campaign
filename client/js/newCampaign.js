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
    const startDate = new Date(campaignStart.value);
    const endDate = new Date(campaignEnd.value);
    const priceValue = parseInt(campaignPrice.value);
    const objectiveValue = parseInt(campaignObjective.value);
    const pricePerDisplay = priceValue / objectiveValue;

    const startDay = startDate.getDate();
    const startMonth = startDate.getMonth() + 1;
    const startYear = startDate.getFullYear();

    const endDay = endDate.getDate();
    const endMonth = endDate.getMonth() + 1;
    const endYear = endDate.getFullYear();

    const formSubmit = {
        "name": nameValue,
        "startDate": {
            "dateDay": startDay,
            "dateMonth": startMonth,
            "dateYear": startYear,
        },
        "endDate": {
            "dateDay": endDay,
            "dateMonth": endMonth,
            "dateYear": endYear,
        },
        "price": priceValue,
        "objective": objectiveValue,
        "pricePerDisplay": pricePerDisplay
    }

    fetch("http://localhost:8080/api/campaign", {
        method: "POST",
        headers: {
            "Content-Type": 'application/json',
        },
        body: JSON.stringify(formSubmit),
    })
    .then(() => {
        window.location.href = "http://localhost:8080"
    })
}

// event when submit 
submit.addEventListener("click", (event) => {
    event.preventDefault();
    newCampaign();
})
