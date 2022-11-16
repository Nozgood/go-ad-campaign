let campaignCard = document.getElementById("campaign__all");
let allCampaigns = [];

fetch("http://localhost:8080/api/campaign/getAll")
        .then((res) => {
            return res.json()
        })
        .then((data) => {
            allCampaigns = data;
            for (let i=0; i < allCampaigns.length; i++) {
                let campaignDiv = document.createElement("div");
                campaignDiv.classList.add("campaign__card")
                let campaignName = document.createElement("h3");
                campaignName.innerHTML = "Nom : " + allCampaigns[i].name;
                let campaignStart = document.createElement("p");
                campaignStart.innerHTML = "Date de début : " + allCampaigns[i].startDate;
                let campaignEnd = document.createElement("p");
                campaignEnd.innerHTML = "Date de fin : " + allCampaigns[i].endDate;
                let campaignPrice = document.createElement("p");
                campaignPrice.innerHTML = "Prix : " +  allCampaigns[i].price;
                let campaignObjective = document.createElement("p");
                campaignObjective.innerHTML = "Objectif d'affichage : " + allCampaigns[i].objective;
                let campaignPricePerDisplay = document.createElement("p");
                campaignPricePerDisplay.innerHTML = "Prix par affichage : " + allCampaigns[i].pricePerDisplay;

                let campaignModify = document.createElement("a");
                campaignModify.setAttribute("href", `http://localhost:8080/one/${allCampaigns[i].name}`)
                campaignModify.innerHTML = "Modifier";
        
                let campaignDelete = document.createElement("button");
                campaignDelete.innerHTML = "Supprimer";

                campaignDiv.appendChild(campaignName);
                campaignDiv.appendChild(campaignStart);
                campaignDiv.appendChild(campaignEnd);
                campaignDiv.appendChild(campaignPrice);
                campaignDiv.appendChild(campaignObjective);
                campaignDiv.appendChild(campaignPricePerDisplay);
                campaignDiv.appendChild(campaignModify);
                campaignCard.appendChild(campaignDiv);
            }
        })
