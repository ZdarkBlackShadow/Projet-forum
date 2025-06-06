// Vos fonctions existantes
function toggleSearch() {
    const container = document.getElementById('searchContainer');
    const currentDisplay = window.getComputedStyle(container).display;
    if (currentDisplay === 'none') {
        container.style.display = 'flex';
    } else {
        container.style.display = 'none';
    }
}

function toggleFriendRequest() {
    const container = document.getElementById('friendRequestSearchContainer');
    const currentDisplay = window.getComputedStyle(container).display;
    if (currentDisplay === 'none') {
        container.style.display = 'flex';
    } else {
        container.style.display = 'none';
    }
}

function toggleAddSearchBar() {
    const container = document.getElementById('addSearchBarContainer');
    const currentDisplay = window.getComputedStyle(container).display;
    if (currentDisplay === 'none') {
        container.style.display = 'block';
    } else {
        container.style.display = 'none';
    }
}

function toggleMessagePopup() {
    const popup = document.getElementById('messagePopup');
    const currentDisplay = window.getComputedStyle(popup).display;
    if (currentDisplay === 'none') {
        popup.style.display = 'block';
    } else {
        popup.style.display = 'none';
    }
}

function deleteConversation() {
    const panel = document.getElementById('conversationPanel');
    panel.innerHTML = ''; 
    panel.style.display = 'none';
    document.getElementById('conversationToggle').style.display = 'none';
}

function removeConversation(convId) {
    const convEl = document.getElementById(convId);
    if (convEl) {
        convEl.remove();
    }
}

let conversationCount = 0;

function acceptMessageRequest() {
    document.getElementById('messagePopup').style.display = 'none';
    conversationCount++;
    const convId = 'conversation-' + conversationCount;
    const convList = document.getElementById('conversationList');
    let convItem = document.createElement('div');
    convItem.id = convId;
    convItem.className = 'conversation';
    convItem.innerHTML = `
        <div class="conversation-avatar">X</div>
        <div class="conversation-name">Utilisateur X</div>
        <button class="delete-conv-btn" onclick="removeConversation('${convId}')">Supprimer</button>
    `;
    convList.appendChild(convItem);
    openConversation(convId);
}

function openConversation(convId) {
    const panel = document.getElementById('conversationPanel');
    panel.style.display = 'block';
    panel.innerHTML = `
        <h4>Conversation avec Utilisateur X</h4>
        <button class="delete-conv-btn" onclick="deleteConversation()">Supprimer conversation</button>
        <!-- Conversation messages go here -->
    `;
}

async function loadEmojis() {
    const panel = document.getElementById('emojiPanel');
    if (panel.innerHTML.trim().length > 0) return; // Already loaded
    try {
        const response = await fetch("https://emoji-api.com/emojis?access_key=YOUR_ACCESS_KEY&limit=20");
        if (response.ok) {
            const emojis = await response.json();
            emojis.forEach(emoji => {
                const span = document.createElement('span');
                span.className = "emoji";
                span.textContent = emoji.character;
                span.onclick = () => insertEmoji(emoji.character);
                panel.appendChild(span);
            });
        } else {
            panel.innerHTML = "Erreur de chargement";
        }
    } catch (err) {
        console.error(err);
        panel.innerHTML = "Erreur de chargement";
    }
}

function toggleEmojiPanel() {
    const panel = document.getElementById('emojiPanel');
    if (window.getComputedStyle(panel).display === "none") {
        panel.style.display = "block";
        loadEmojis();
    } else {
        panel.style.display = "none";
    }
}

function insertEmoji(emoji) {
    const messageInput = document.getElementById('messageInput');
    messageInput.value += emoji;
    document.getElementById('emojiPanel').style.display = 'none';
}