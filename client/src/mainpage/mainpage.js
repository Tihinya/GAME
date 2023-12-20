export function showWriteName() {
  let playButton = document.getElementById("play-button");
  let nameContainer = document.getElementById("name-container");
  // Hide the play button
  playButton.style.display = "none";

  // Show the write name div
  nameContainer.style.display = "flex";
}

document.getElementById("play-button").addEventListener("click", showWriteName);
