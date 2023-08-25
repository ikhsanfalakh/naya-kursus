class TableManager {
  constructor(tableId) {
    this.table = document.getElementById(tableId);
  }

  addRow(data) {
    const newRow = this.table.insertRow();
    for (const value of data) {
      const cell = newRow.insertCell();
      cell.textContent = value;
    }
  }
}
  
const tableManager = new TableManager("dataTable");
const addButton = document.getElementById("addButton");

addButton.addEventListener("click", (event) => {
  event.preventDefault(); // Mencegah perilaku bawaan formulir

  const namaBarang = document.getElementById("namaBarang").value;
  const jumlahBarang = document.getElementById("jumlahBarang").value;
  const harga = document.getElementById("harga").value;

  if (namaBarang.trim() !== "" && jumlahBarang.trim() !== "" && harga.trim() !== "") {
    const data = [namaBarang, jumlahBarang, harga];
    tableManager.addRow(data);

    document.getElementById("namaBarang").value = "";
    document.getElementById("jumlahBarang").value = "";
    document.getElementById("harga").value = "";
  }
});
  