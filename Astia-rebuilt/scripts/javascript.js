// Selecting template and preview Image
const btn = document.querySelector('#flu-btn');        
const radioButtons = document.querySelectorAll('input[name="flexRadioDefault"]');

btn.addEventListener("click", () => {
    let selectedTemplate;
    for (const radioButton of radioButtons) {
        if (radioButton.checked) {
            selectedTemplate = radioButton.value;
            break;
        }
    }
    // show the output:
    output.innerText = selectedTemplate ? `You selected ${selectedTemplate} Template. Go to Next tab to set Target and launch Attack` : `You haven't selected any template`;
});

function changeSrc() {
    if (document.getElementById("flexRadioDefault1").checked) {
    document.getElementById("preview-img").src = "images/image1.png";
    }
    else if (document.getElementById("flexRadioDefault2").checked) { 
    document.getElementById("preview-img").src = "images/image2.png";
    }
    else if (document.getElementById("flexRadioDefault3").checked) {
        document.getElementById("preview-img").src = "images/signin.jpg";
        }
    else if (document.getElementById("flexRadioDefault4").checked) {
        document.getElementById("preview-img").src = "images/lottery.png";
        }
    else if (document.getElementById("flexRadioDefault5").checked) {
        document.getElementById("preview-img").src = "images/flu.png";
        }
}