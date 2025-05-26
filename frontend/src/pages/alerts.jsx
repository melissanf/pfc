import React, { useState, useEffect } from 'react'; // ✅ Ajout de useEffect
import PageWrapper from '../components/PageWrapper';
import Sidebar from '../components/Sidebar';
import './alerts.css';
import { useNavigate } from 'react-router-dom'; // ✅ Si tu utilises `navigate`

export default function Alerts() {
  const [alerts, setAlerts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate(); // ✅ Si tu veux rediriger

  useEffect(() => {
    const userString = localStorage.getItem("user");

    let userId = null;
    if (userString) {
      try {
        const user = JSON.parse(userString);
        userId = user.ID;
        const token = localStorage.getItem('authToken') || 
              localStorage.getItem('token') || 
              localStorage.getItem('accessToken');

        fetch(`http://localhost:8000/Enseignant/notifications?user_id=${userId}`,{
          method: "GET",
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` // Assurez-vous que le token est correct
          } 
        })
          .then((res) => {
            if (!res.ok) throw new Error('Erreur serveur');
            return res.json();
          })
          .then((data) => {
            setAlerts(data.notifications || []);
            setLoading(false);
          })
          .catch((err) => {
            console.error(err);
            setError("Impossible de charger les notifications");
            setLoading(false);
          });

      } catch (error) {
        console.error("Erreur lors du parsing des données utilisateur:", error);
        navigate("/dashboardtec");
      }
    } else {
      console.warn("Aucun utilisateur trouvé dans le local storage.");
      navigate("/login");
    }
  }, [navigate]);

  const handleDelete = (id) => {
    const confirmation = window.confirm('Voulez-vous vraiment supprimer cette alerte ?');
    if (confirmation) {
      setAlerts(prevAlerts => prevAlerts.filter(alert => alert.id !== id));
    }
  };

  return (
    <PageWrapper>
      <div className="alerts-container">
        <Sidebar />
        <main className="alert-main">
          <h2 className="title">Alertes et Notifications</h2>

          {loading && <p>Chargement...</p>}
          {error && <p className="error">{error}</p>}

          {!loading && !error && alerts.map((alert) => (
            <div key={alert.id} className={`alert-card ${alert.type || 'blue'}`}>
              <div className="alert-content">
                <h3>{alert.titre}</h3>
                <p>{alert.message}</p>
                <div className="meta">
                  <span>{new Date(alert.created_at).toLocaleDateString()}</span>
                  {!alert.is_read && <span className="unread">🟢 Non lu</span>}
                </div>
              </div>
              <button
                className="delete-button"
                onClick={() => handleDelete(alert.id)}
              >
                Supprimer
              </button>
            </div>
          ))}
        </main>
      </div>
    </PageWrapper>
  );
}
