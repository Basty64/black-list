// Placeholder - реальная логика должна быть в бекенде
const blacklistData = [
    {
        "fio": "Иванов Иван Иванович",
        "passport_number": "1234567890",
        "passport_series": "1234",
        "previous_passport_number": null,
        "previous_passport_series": null,
        "company_employee": true,
        "organization": null,
        "reason": "Кража имущества",
        "date": "2023-03-15"
    },
    // ... еще записи
];

// Функции для работы с данными
function getBlacklistEntries() {
    return blacklistData;
}

function addBlacklistEntry(newEntry) {
    blacklistData.push(newEntry);
    // Запись в бд
}

function updateBlacklistEntry(entryId, updatedEntry) {
    // Обновление записи в бд
}

function deleteBlacklistEntry(entryId) {
    // Удаление записи из бд
}

// Функции для отображения таблицы
function renderBlacklistTable() {
    const tableBody = document.getElementById('blacklist-table-body');
    tableBody.innerHTML = ''; // Очистка таблицы

    const entries = getBlacklistEntries();

    entries.forEach(entry => {
        const row = tableBody.insertRow();
        const fioCell = row.insertCell();
        const passportNumberCell = row.insertCell();
        const passportSeriesCell = row.insertCell();
        const companyEmployeeCell = row.insertCell();
        const organizationCell = row.insertCell();
        const reasonCell = row.insertCell();
        const dateCell = row.insertCell();
        const actionsCell = row.insertCell();

        fioCell.textContent = entry.fio;
        passportNumberCell.textContent = entry.passport_number;
        passportSeriesCell.textContent = entry.passport_series;
        companyEmployeeCell.textContent = entry.company_employee ? 'Да' : 'Нет';
        organizationCell.textContent = entry.organization || '—';
        reasonCell.textContent = entry.reason;
        dateCell.textContent = entry.date;
        actionsCell.innerHTML = `
      <button class="edit-button" data-entry-id="${entry.id}">Редактировать</button>
      <button class="delete-button" data-entry-id="${entry.id}">Удалить</button>
    `;
    });
}

// Функции для управления модальными окнами
function openModal(title, entry = null) {
    const modal = document.getElementById('modal');
    const modalTitle = document.getElementById('modal-title');
    const modalForm = document.getElementById('modal-form');
    const modalSubmitButton = document.getElementById('modal-submit-button');

    modalTitle.textContent = title;
    modalForm.reset(); // Очистка формы

    if (entry) {
        // Заполнение формы данными из записи
        document.getElementById('modal-entry-id').value = entry.id;
        document.getElementById('modal-fio').value = entry.fio;
        // ... заполнение остальных полей
    } else {
        document.getElementById('modal-entry-id').value = '';
    }

    modal.classList.remove('hidden');
}

function closeModal() {
    const modal = document.getElementById('modal');
    modal.classList.add('hidden');
}

// Обработка событий
const addEntryButton = document.getElementById('add-entry-button');
addEntryButton.addEventListener('click', () => openModal('Добавить запись'));

const searchButton = document.getElementById('search-button');
searchButton.addEventListener('click', () => {
    // Логика поиска
});

// ... другие обработчики событий

// Инициализация приложения
renderBlacklistTable();