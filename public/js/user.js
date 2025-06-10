document.addEventListener('DOMContentLoaded', () => {
    console.log('User profile settings page loaded.');

    const avatarPreview = document.getElementById('avatarPreview');
    if (avatarPreview) {
        // Event listener for when the image successfully loads
        avatarPreview.addEventListener('load', function() {
            setBannerColorFromAvatar(this); // 'this' refers to avatarPreview
        });

        // Event listener for an error during image loading
        avatarPreview.addEventListener('error', function() {
            console.error('Avatar image failed to load. Setting default banner color.');
            const banner = document.querySelector('.profile-banner-preview');
            const bannerImage = document.getElementById('bannerPreview');
            if (banner) {
                banner.style.backgroundColor = '#5865f2'; // Default Discord blue
            }
            if (bannerImage) {
                bannerImage.style.display = 'block'; // Ensure default banner image is visible
            }
        });

        // If the image is already loaded from cache when the script runs
        if (avatarPreview.complete && avatarPreview.naturalHeight !== 0) {
            setBannerColorFromAvatar(avatarPreview);
        } else if (avatarPreview.complete && avatarPreview.naturalHeight === 0 && avatarPreview.src) {
            // Image is 'complete' but has no dimensions, could be an error with the src or an empty image
            console.warn('Avatar image reported as complete but has no dimensions. Setting default banner color.');
            const banner = document.querySelector('.profile-banner-preview');
            const bannerImage = document.getElementById('bannerPreview');
            if (banner) {
                banner.style.backgroundColor = '#5865f2';
            }
            if (bannerImage) {
                bannerImage.style.display = 'block';
            }
        }
    }

    // Example: Load data if available
    // const userData = {
    //     username: "VotrePseudo",
    //     tag: "1234",
    //     customStatus: "Chilling...",
    //     aboutMe: "Loves coding and gaming!",
    //     memberSince: "15 Fév. 2021",
    //     avatarUrl: "/public/img/default_image.png", // This would trigger the 'load' event
    //     bannerUrl: "/public/img/default_banner.jpg" // Banner image is handled separately or hidden
    // };
    // loadProfileData(userData); // Call this if you load profile data dynamically
});

function setBannerColorFromAvatar(imgElement) {
    const banner = document.querySelector('.profile-banner-preview');
    const bannerImage = document.getElementById('bannerPreview'); // The actual <img> for the banner
    if (!banner) return;

    // Ensure the image element is valid and loaded
    if (!imgElement || !imgElement.complete || imgElement.naturalWidth === 0 || imgElement.naturalHeight === 0) {
        console.warn('Avatar image not ready for color extraction. Using default banner color.');
        banner.style.backgroundColor = '#5865f2';
        if (bannerImage) bannerImage.style.display = 'block';
        return;
    }

    try {
        const canvas = document.createElement('canvas');
        const context = canvas.getContext('2d');

        if (!context) {
            console.warn('Canvas 2D context not supported. Using default banner color.');
            banner.style.backgroundColor = '#5865f2';
            if (bannerImage) bannerImage.style.display = 'block';
            return;
        }

        canvas.width = imgElement.naturalWidth;
        canvas.height = imgElement.naturalHeight;
        context.drawImage(imgElement, 0, 0, canvas.width, canvas.height);

        const imageData = context.getImageData(0, 0, canvas.width, canvas.height).data;
        let r = 0, g = 0, b = 0;
        let pixelCount = 0;

        // Iterate over pixels to calculate average color
        // Consider only mostly opaque pixels
        for (let i = 0; i < imageData.length; i += 4) {
            const alpha = imageData[i + 3];
            if (alpha > 200) { // Only consider pixels that are substantially opaque
                r += imageData[i];
                g += imageData[i + 1];
                b += imageData[i + 2];
                pixelCount++;
            }
        }

        if (pixelCount === 0) {
            // No suitable pixels found (e.g., fully transparent image)
            banner.style.backgroundColor = '#5865f2'; // Default color
            if (bannerImage) bannerImage.style.display = 'block';
            return;
        }

        r = Math.floor(r / pixelCount);
        g = Math.floor(g / pixelCount);
        b = Math.floor(b / pixelCount);
        
        // Basic adjustment to avoid very dark or very light muddy colors
        const brightness = (r * 0.299 + g * 0.587 + b * 0.114); // Luminance
        if (brightness < 50 && (r < 80 || g < 80 || b < 80) ) { // If too dark, lighten
            r = Math.min(255, r + 30);
            g = Math.min(255, g + 30);
            b = Math.min(255, b + 30);
        } else if (brightness > 200) { // If too light, darken
            r = Math.max(0, r - 30);
            g = Math.max(0, g - 30);
            b = Math.max(0, b - 30);
        }

        banner.style.backgroundColor = `rgb(${r}, ${g}, ${b})`;
        if (bannerImage) {
            bannerImage.style.display = 'none'; // Hide the default banner image to show the color
        }

    } catch (e) {
        console.error('Error extracting color from avatar:', e);
        banner.style.backgroundColor = '#5865f2'; // Fallback to default color
        if (bannerImage) bannerImage.style.display = 'block'; // Show default banner image
    }
}

function loadProfileData(data) {
    const usernameTagPreview = document.getElementById('usernameTagPreview');
    const customStatusInput = document.getElementById('customStatusInput');
    const aboutMeTextarea = document.getElementById('aboutMeTextarea');
    const memberSinceDate = document.getElementById('memberSinceDate');
    const avatarPreview = document.getElementById('avatarPreview');
    // const bannerPreviewImg = document.getElementById('bannerPreview'); // The <img> for banner

    if (data.username && data.tag && usernameTagPreview) {
        usernameTagPreview.textContent = `${data.username}#${data.tag}`;
    }
    if (data.customStatus && customStatusInput) {
        customStatusInput.value = data.customStatus;
    }
    if (data.aboutMe && aboutMeTextarea) {
        aboutMeTextarea.value = data.aboutMe;
    }
    if (data.memberSince && memberSinceDate) {
        memberSinceDate.textContent = data.memberSince;
    }
    if (data.avatarUrl && avatarPreview) {
        avatarPreview.src = data.avatarUrl; // This will trigger the 'load' or 'error' event
    }
    // If you have a specific banner image URL and want to use it instead of color
    // if (data.bannerUrl && bannerPreviewImg) {
    //     const bannerDiv = document.querySelector('.profile-banner-preview');
    //     bannerPreviewImg.src = data.bannerUrl;
    //     bannerPreviewImg.style.display = 'block';
    //     bannerDiv.style.backgroundColor = ''; // Clear background color if image is set
    // }
}

function clearCustomStatus() {
    const statusInput = document.getElementById('customStatusInput');
    if (statusInput) {
        statusInput.value = '';
    }
}

function saveProfileChanges() {
    const customStatus = document.getElementById('customStatusInput').value;
    const aboutMe = document.getElementById('aboutMeTextarea').value;

    console.log('Saving profile changes:');
    console.log('Custom Status:', customStatus);
    console.log('About Me:', aboutMe);

    alert('Modifications (simulées) sauvegardées !');
}

// Placeholder for image upload handling (more complex)
// document.getElementById('avatarUpload').addEventListener('change', function(event) {
//     const file = event.target.files[0];
//     if (file) {
//         const reader = new FileReader();
//         reader.onload = function(e) {
//             document.getElementById('avatarPreview').src = e.target.result;
//         }
//         reader.readAsDataURL(file);
//     }
// });
