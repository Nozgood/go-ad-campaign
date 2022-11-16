let url = new URL(window.location.href);
let campaignName = url.href.split(':8080/')[1];

let campaignDiv = document.getElementById("oneCampaign");

fetch("http://localhost:8080/api/campaign/getByName/" + campaignName)
    .then((res) => {return res.json()})
    .then((data) => {
        let campaignName = document.createElement("h3");
        campaignName.innerHTML = "Nom : " + data.name;
        let campaignStart = document.createElement("p");
        campaignStart.innerHTML = "Date de début : " + data.startDate;
        let campaignEnd = document.createElement("p");
        campaignEnd.innerHTML = "Date de fin : " + data.endDate;
        let campaignPrice = document.createElement("p");
        campaignPrice.innerHTML = "Prix : " +  data.price;
        let campaignObjective = document.createElement("p");
        campaignObjective.innerHTML = "Objectif d'affichage : " + data.objective;
        let campaignPricePerDisplay = document.createElement("p");
        campaignPricePerDisplay.innerHTML = "Prix par affichage : " + data.pricePerDisplay;
        
        campaignDiv.appendChild(campaignName);
        campaignDiv.appendChild(campaignStart);
        campaignDiv.appendChild(campaignEnd);
        campaignDiv.appendChild(campaignPrice);
        campaignDiv.appendChild(campaignObjective);
        campaignDiv.appendChild(campaignPricePerDisplay);
    });