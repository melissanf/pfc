import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./WishList.css";
import { FaArrowLeft, FaSave, FaPlus } from "react-icons/fa";

const TeachingTypeSelector = ({ selectedTypes = [], onChange }) => {
  const types = ["Cours", "TD", "TP"];

  const toggleType = (type) => {
    if (selectedTypes.includes(type)) {
      onChange(selectedTypes.filter((t) => t !== type));
    } else {
      onChange([...selectedTypes, type]);
    }
  };

  return (
    <div className="form-group-full">
      <label className="section-title">Type d’enseignement</label>
      <div className="toggle-buttons-row">
        {types.map((type) => (
          <button
            key={type}
            className={`toggle-button ${selectedTypes.includes(type) ? "active" : ""}`}
            onClick={() => toggleType(type)}
            type="button"
          >
            {type}
          </button>
        ))}
      </div>
    </div>
  );
};

const WishItem = ({ index, wish, onUpdate, onDelete, isExtra = false }) => {
  const handleModuleChange = (e) => {
    onUpdate(index, { ...wish, module: e.target.value });
  };

  const handleNiveauChange = (e) => {
    onUpdate(index, { ...wish, niveau: e.target.value });
  };

  const handleTypesChange = (types) => {
    onUpdate(index, { ...wish, teachingTypes: types });
  };

  const handleDelete = () => {
    if (window.confirm("Voulez-vous vraiment supprimer ce vœu ?")) {
      onDelete(index);
    }
  };

  return (
    <div className="form-card">
      <div className="WishList-header">
        <h3>{isExtra ? "Heures Supplémentaires" : `Vœu ${index + 1}`}</h3>
        <button className="delete-btn" onClick={handleDelete}>
          Supprimer
        </button>
      </div>

      <div className="form-row">
        <div className="form-group">
          <label>Nom du module</label>
          <input
            type="text"
            placeholder="Entrez le nom du module"
            value={wish.module}
            onChange={handleModuleChange}
          />
        </div>
      </div>

      <div className="form-row">
        <div className="form-group">
          <label>Niveau</label>
          <input
            type="text"
            placeholder="Ex : L3, M1, etc."
            value={wish.niveau || ""}
            onChange={handleNiveauChange}
          />
        </div>
      </div>

      <TeachingTypeSelector
        selectedTypes={wish.teachingTypes}
        onChange={handleTypesChange}
      />
    </div>
  );
};

function WishList() {
  const [wishList, setWishList] = useState([]);
  const [extraWish, setExtraWish] = useState(null);
  const [saveMessage, setSaveMessage] = useState("");
  const navigate = useNavigate();

  const handleBack = () => {
    navigate("/dashboardtec"); // Modifie la route si besoin
  };

  useEffect(() => {
    fetch("http://localhost:8000/Enseignant/fiche-de-voeux")
      .then((res) => {
        if (!res.ok) throw new Error("Erreur réseau");
        return res.json();
      })
      .then((data) => {
        const wishes = data.map((item) => ({
          module: item.module_name,
          niveau: item.niveau_name,
          teachingTypes: [
            item.cour && "Cours",
            item.td && "TD",
            item.tp && "TP",
          ].filter(Boolean),
          isExtra: item.hr === true,
        }));

        const mainWishes = wishes.filter(w => !w.isExtra).slice(0, 3);
        const extra = wishes.find(w => w.isExtra);

        setWishList(mainWishes);
        setExtraWish(extra || null);
      })
      .catch((err) => {
        console.error("Erreur lors du chargement des vœux depuis l'API:", err);
      });
  }, []);

  const addWish = () => {
    if (wishList.length >= 3) return;
    setWishList([...wishList, { module: "", niveau: "", teachingTypes: [] }]);
  };

  const updateWish = (index, updatedWish) => {
    const newList = [...wishList];
    newList[index] = updatedWish;
    setWishList(newList);
  };

  const deleteWish = (index) => {
    const newList = wishList.filter((_, i) => i !== index);
    setWishList(newList);
  };

  const addExtraWish = () => {
    setExtraWish({ module: "", niveau: "", teachingTypes: [] });
  };

  const updateExtraWish = (_, updatedWish) => {
    setExtraWish(updatedWish);
  };

const saveWishes = async () => {
  const allWishes = [...wishList];
  if (extraWish) allWishes.push({ ...extraWish, hr: true });

  // Validation
  const hasEmptyFields = allWishes.some(
    (wish) =>
      !wish.module?.trim() ||
      !wish.niveau?.trim() ||
      wish.teachingTypes.length === 0
  );
  
  if (hasEmptyFields) {
    setSaveMessage("❗ Veuillez remplir tous les champs avant d'enregistrer.");
    setTimeout(() => setSaveMessage(""), 3000);
    return;
  }

  // Format pour API
  const payload = allWishes.map((wish) => ({
    module_name: wish.module,
    niveau_name: wish.niveau,
    cour: wish.teachingTypes.includes("Cours"),
    td: wish.teachingTypes.includes("TD"),
    tp: wish.teachingTypes.includes("TP"),
    hr: wish.hr || false,
  }));

  try {
    const token = localStorage.getItem('authToken') || 
                  localStorage.getItem('token') || 
                  localStorage.getItem('accessToken');
    
    const userRole = localStorage.getItem('userRole');
    const userId = localStorage.getItem('userId') || localStorage.getItem('user_id');

    if (!token) {
      setSaveMessage("❌ Token d'authentification manquant. Veuillez vous reconnecter.");
      return;
    }

    console.log("User role:", userRole);
    console.log("User ID:", userId);
    console.log("Payload:", JSON.stringify(payload, null, 2));

    // Headers enrichis avec plus d'informations
    const headers = {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`,
    };

    // Ajouter des headers supplémentaires si disponibles
    if (userRole) {
      headers["X-User-Role"] = userRole;
    }
    if (userId) {
      headers["X-User-ID"] = userId;
    }

    console.log("Headers envoyés:", headers);

    const response = await fetch("http://localhost:8000/Enseignant/fiche-de-voeux", {
      method: "POST",
      headers: headers,
      body: JSON.stringify(payload),
    });

    console.log("Response status:", response.status);
    console.log("Response headers:", Object.fromEntries(response.headers.entries()));

    if (!response.ok) {
      const errorText = await response.text();
      console.log("Error response:", errorText);
      
      let errorMessage;
      try {
        const errorData = JSON.parse(errorText);
        errorMessage = errorData.message || errorData.detail || errorData.error;
      } catch (e) {
        errorMessage = errorText;
      }

      if (response.status === 401) {
        setSaveMessage("❌ Token expiré. Veuillez vous reconnecter.");
      } else if (response.status === 403) {
        setSaveMessage(`❌ Accès refusé: ${errorMessage || 'Permissions insuffisantes'}`);
      } else {
        setSaveMessage(`❌ Erreur ${response.status}: ${errorMessage}`);
      }
      return;
    }

    const responseData = await response.json();
    console.log("Success response:", responseData);

    // Succès
    setWishList([]);
    setExtraWish(null);
    setSaveMessage("✅ Vœux enregistrés avec succès !");
    setTimeout(() => setSaveMessage(""), 3000);

  } catch (error) {
    console.error("Network error:", error);
    setSaveMessage(`❌ Erreur réseau: ${error.message}`);
    setTimeout(() => setSaveMessage(""), 5000);
  }
};
  return (
    <div className="wishlist-container">
      <button className="back-button" onClick={handleBack}>
        <FaArrowLeft className="icon" size={12} /> Retour
      </button>
      <h1 className="title">Liste de Vœux</h1>
      <p className="subtitle">
        Gérez vos souhaits d’enseignement pour l’année à venir
      </p>

      {wishList.map((wish, index) => (
        <WishItem
          key={index}
          index={index}
          wish={wish}
          onUpdate={updateWish}
          onDelete={deleteWish}
        />
      ))}

      {extraWish && (
        <WishItem
          index={3}
          wish={extraWish}
          onUpdate={updateExtraWish}
          onDelete={() => setExtraWish(null)}
          isExtra={true}
        />
      )}

      <div className="buttons-container">
        {wishList.length < 3 && (
          <button className="add-button" onClick={addWish}>
            <FaPlus className="icon" size={12} /> Nouveau vœu
          </button>
        )}

        {wishList.length > 0 && (
          <button className="save-button" onClick={saveWishes}>
            <FaSave className="icon" size={12} /> Enregistrer les vœux
          </button>
        )}

        {wishList.length === 3 && !extraWish && (
          <button className="extra-button" onClick={addExtraWish}>
            <FaPlus className="icon" size={12} /> Heures supplémentaires
          </button>
        )}
      </div>

      {wishList.length >= 3 && (
        <p className="limit-message">
          Vous avez atteint le nombre maximum de 3 vœux.
        </p>
      )}

      {saveMessage && <p className="save-message">{saveMessage}</p>}
    </div>
  );
}

export default WishList;
