const filterBtn = document.querySelector(".filter-btn")
const filterHolder = document.querySelector(".filters-holder")
const filterCancelBtn = document.querySelector(".filters-exit-btn")
const showFilters = () => {
    filterHolder.classList.add("show")
    document.body.style.overflow = "hidden"
}
const hideFilters = () => {
    filterHolder.classList.remove("show")
    document.body.style.overflow = ""
}

filterBtn.addEventListener("click", showFilters)
filterCancelBtn.addEventListener("click", hideFilters)

// Range functionality
const gap = 3
const creationRangeInputs = document.querySelectorAll(".creation-filter .range-input input")
const creationRange = document.querySelector(".creation-filter .slider .progress")
const creationResultRange = document.querySelector(".creation-filter .filter-result")

const firstAlbumRangeInputs = document.querySelectorAll(".first-album-filter .range-input input")
const firstAlbumRange = document.querySelector(".first-album-filter .slider .progress")
const firstAlbumResultRange = document.querySelector(".first-album-filter .filter-result")

const rangeFunc = (e, rangeInputs, resultRange, range, gap) => {
    let minValue = parseInt(rangeInputs[0].value),
    maxValue = parseInt(rangeInputs[1].value)

    if (e && (maxValue - minValue) < gap) {
        if (e.target.classList.contains("range-min")) {
            rangeInputs[0].value = maxValue - gap
            minValue = parseInt(rangeInputs[0].value)
        } else {
            rangeInputs[1].value = minValue + gap
            maxValue = parseInt(rangeInputs[1].value)
        }
    } 

    resultRange.textContent = `${minValue} - ${maxValue}`

    const gapValue = rangeInputs[0].max - rangeInputs[0].min
    const minGabValue = ((rangeInputs[0].max - rangeInputs[0].value) / gapValue) * 100
    const maxGabValue = 100-((rangeInputs[0].max - rangeInputs[1].value) / gapValue) *100

    range.style.left = 100-minGabValue+"%"
    range.style.right = 100-maxGabValue+"%"
}

rangeFunc(null, creationRangeInputs, creationResultRange, creationRange, gap)
rangeFunc(null, firstAlbumRangeInputs, firstAlbumResultRange, firstAlbumRange, gap)

creationRangeInputs.forEach(input => {
    input.addEventListener("input", (e) => {rangeFunc(e, creationRangeInputs, creationResultRange, creationRange, gap)})
});

firstAlbumRangeInputs.forEach(input => {
    input.addEventListener("input", (e) => {rangeFunc(e, firstAlbumRangeInputs, firstAlbumResultRange, firstAlbumRange, gap)})
})