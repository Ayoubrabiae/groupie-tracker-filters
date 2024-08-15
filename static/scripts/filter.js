const filterBtn = document.querySelector(".filter-btn")
const filterHolder = document.querySelector(".filter-holder")

filterBtn.addEventListener("click", () => {
    console.log("Hello")
})

// Range functionality
const rangeInputs = document.querySelectorAll(".range-input input")
const range = document.querySelector(".slider .progress")
const resultRange = document.querySelector(".filter-item-holder .filter-result")
const gap = 3

rangeInputs.forEach(input => {
    input.addEventListener("input", (e) => {
        let minValue = parseInt(rangeInputs[0].value),
        maxValue = parseInt(rangeInputs[1].value)

        if ((maxValue - minValue) < gap) {
            if (e.target.classList.contains("range-min")) {
                rangeInputs[0].value = maxValue - gap
            } else {
                rangeInputs[1].value = minValue + gap
            }
        } else {
            resultRange.textContent = `${minValue} - ${maxValue}`
    
            const gapValue = rangeInputs[0].max - rangeInputs[0].min
            const minGabValue = ((rangeInputs[0].max - rangeInputs[0].value) / gapValue) * 100
            const maxGabValue = 100-((rangeInputs[0].max - rangeInputs[1].value) / gapValue) *100
    
            range.style.left = 100-minGabValue+"%"
            range.style.right = 100-maxGabValue+"%"
    
            console.log({gapValue, minGabValue, maxGabValue})
        }

    })
});