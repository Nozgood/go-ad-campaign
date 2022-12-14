let campaignCard = document.getElementById("campaign__all");
let allCampaigns = [];
let allCampaignsDate = [];
let allCampaignsName = [];
const filterButton = document.getElementById("filterButton");
const cancelFilterButton = document.getElementById("filterCancelButton");
const searchBar = document.getElementById("searchBar");
const searchButton = document.getElementById("searchButton");
const searchResults = document.getElementById("searchResults");

const filterByDate = (allCampaignsDate) => {
    const dayDate = new Date();
    allCampaignsDate.forEach((campaignObject) => {
        if (dayDate > campaignObject.end || dayDate < campaignObject.start) {
            const campaignToRemove = document.getElementById(campaignObject.name);
            campaignToRemove.remove();
        }
    })
}

fetch("http://localhost:8080/api/campaign")
        .then((res) => {
            return res.json()
        })
        .then((data) => {
            console.log(data);
            allCampaigns = data;
            for (let i=0; i < allCampaigns.length; i++) {
                let campaignDiv = document.createElement("div");
                campaignDiv.classList.add("campaign__card");
                campaignDiv.setAttribute("id", allCampaigns[i].name);
                let campaignName = document.createElement("h3");
                campaignName.innerHTML = "Nom : " + allCampaigns[i].name;
                let campaignStart = document.createElement("p");
                campaignStart.innerHTML = "Date de début : " + allCampaigns[i].startDate.dateDay + "/" + allCampaigns[i].startDate.dateMonth + "/" +allCampaigns[i].startDate.dateYear;
                let campaignEnd = document.createElement("p");
                campaignEnd.innerHTML = "Date de fin : " + allCampaigns[i].endDate.dateDay + "/" + allCampaigns[i].endDate.dateMonth + "/" +allCampaigns[i].endDate.dateYear;
                let campaignPrice = document.createElement("p");
                campaignPrice.innerHTML = "Prix (€) : " +  allCampaigns[i].price;
                let campaignObjective = document.createElement("p");
                campaignObjective.innerHTML = "Objectif d'affichage : " + allCampaigns[i].objective;
                let campaignPricePerDisplay = document.createElement("p");
                campaignPricePerDisplay.innerHTML = "Prix par affichage (€) : " + allCampaigns[i].pricePerDisplay;

                let campaignModify = document.createElement("a");
                campaignModify.setAttribute("href", `http://localhost:8080/${allCampaigns[i].name}`)
                campaignModify.innerHTML = "Afficher";

                campaignDiv.appendChild(campaignName);
                campaignDiv.appendChild(campaignStart);
                campaignDiv.appendChild(campaignEnd);
                campaignDiv.appendChild(campaignPrice);
                campaignDiv.appendChild(campaignObjective);
                campaignDiv.appendChild(campaignPricePerDisplay);
                campaignDiv.appendChild(campaignModify);
                campaignCard.appendChild(campaignDiv);

                const startDate = new Date(allCampaigns[i].startDate.dateYear, (allCampaigns[i].startDate.dateMonth -1), allCampaigns[i].startDate.dateDay);
                const endDate = new Date(allCampaigns[i].endDate.dateYear, (allCampaigns[i].endDate.dateMonth -1), allCampaigns[i].endDate.dateDay);
                const campaignObject = {
                    "name": allCampaigns[i].name,
                    "start": startDate,
                    "end": endDate,
                }
                allCampaignsDate.push(campaignObject);
                allCampaignsName.push(allCampaigns[i].name);
            }
        })

// event to listen to filter campaigns 
filterButton.addEventListener("click", () => {
    filterByDate(allCampaignsDate);
})

// cancel the filter
cancelFilterButton.addEventListener("click", () => {
    window.location.reload();
})

// Search Bar
// register the value of the search input
let searchValue;

searchBar.addEventListener("change", (e) => {
    searchValue = e.target.value;
    return searchValue;
})

searchButton.addEventListener("click", () => {

    const linksAlready = document.querySelectorAll("div#searchResults > a"); 
    if (linksAlready.length > 0) {
        linksAlready.forEach((link) => {
            link.remove();
        })
    }

    allCampaignsName.filter((value) => {
        return (
            value.toLowerCase().includes(searchValue.toLowerCase())
        )
    }).map((value) => {
            const searchLink = document.createElement("a");
            searchLink.setAttribute("href", "http://localhost:8080/" + value);
            searchLink.innerHTML = value;
            searchResults.appendChild(searchLink);
    })
})