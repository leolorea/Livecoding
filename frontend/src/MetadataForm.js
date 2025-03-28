import React, { useState } from "react";
import axios from "axios";

const MetadataForm = () => {
  const [file, setFile] = useState(null);
  const [metadata, setMetadata] = useState([]);
  const [loading, setLoading] = useState(true); // Novo estado para verificar se os dados estão carregando

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const token = "ds%IOF2e2!D&@gd#dsa#hulwG(*d(@98d29`*d@Y*)";

  const handleFileUpload = async (e) => {
    e.preventDefault();

    if (!file) {
      alert("Please, select a file.");
      return;
    }

    const formData = new FormData();
    formData.append("data", file);

    try {
      await axios.post("http://localhost:8080/files", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
          "Authorization": `Bearer ${token}`,
        },
      });
      alert("File uploaded successfully!");
      fetchMetadata();
    } catch (err) {
      console.error("Error uploading file:", err);
      alert("Error uploading file");
    }
  };

  const fetchMetadata = async () => {
    setLoading(true);  // Marcar como carregando ao buscar os dados
    try {
      const response = await axios.get("http://localhost:8080/files", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      console.log("Fetched metadata:", response.data); // Verifique o que o backend está retornando
      if (Array.isArray(response.data)) {
        setMetadata(response.data);
      } else {
        console.error("Expected an array but got:", response.data);
      }
    } catch (err) {
      console.error("Error fetching records:", err);
    } finally {
      setLoading(false);  // Marcar como terminado de carregar
    }
  };

  // Carregar os registros de metadata ao montar o componente
  React.useEffect(() => {
    fetchMetadata();
  }, []);

  return (
    <div>
      <h1>File Upload and Metadata Listing</h1>

      <form onSubmit={handleFileUpload}>   
        <div>
          <label htmlFor="file">Select a file:</label>
          <input
            type="file"
            id="file"
            onChange={handleFileChange}
            required
          />
        </div>

        <button type="submit">Upload File</button>
      </form>

      <h2>Metadata Records</h2>
      
      {loading ? (
        <p>Loading...</p>  // Exibe uma mensagem enquanto os dados são carregados
      ) : (
        <ul>
          {metadata.length === 0 ? (
            <li>No metadata found.</li>  // Exibe uma mensagem se não houver dados
          ) : (
            metadata.map((item) => (
              <li key={item.id}>
                <strong>{item.name}</strong> - <a href={`data:application/octet-stream;base64,${item.data}`} download={item.name}>Download</a>
              </li>
            ))
          )}
        </ul>
      )}
    </div>
  );
};

export default MetadataForm;
