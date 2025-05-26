import React, { useEffect, useState } from "react";
import { Routes, Route, useLocation, Navigate } from "react-router-dom";
import { AnimatePresence, motion } from "framer-motion";

// Pages d'authentification
import Signup from "./components/signup/signup";
import LoginPage from "./components/login/login";

// Pages principales
import Parametre from "./pages/Parametre";
import Alerts from "./pages/alerts";
import ModuleManagement from "./pages/ModuleManagement";
import Commentaires from "./pages/Commentaires";
import WishList from "./pages/WishList";
import Profil from "./pages/Profil";
import TeacherTableManagment from "./pages/TeacherTableManagment";
import OrganigrammePage from "./pages/OrganigrammePage";
import Dashboardorga from "./components/Dashboardorga"; // Chef / Staff
import Dashboardtec from "./components/Dashboardtec"; // Enseignant
import CodesPage from "./pages/codesPage";

// Wrapper animation page transitions
const PageWrapper = ({ children }) => (
  <motion.div
    initial={{ opacity: 0 }}
    animate={{ opacity: 1 }}
    exit={{ opacity: 0 }}
    transition={{ duration: 0.4, ease: "easeInOut" }}
  >
    {children}
  </motion.div>
);

const App = () => {
  const location = useLocation();
  const [role, setRole] = useState("");

  useEffect(() => {
    const userString = localStorage.getItem("user");
    if (userString) {
      try {
        const user = JSON.parse(userString);
        const role = user.role;
        setRole(role);
        console.log("Rôle chargé:", role); // Debug
      } catch (error) {
        console.error("Erreur parsing user:", error);
      }
    } else {
      console.warn("Aucun utilisateur trouvé dans le local storage.");
    }
  }, []);

  const handleRoleChange = (newRole) => {
    const userString = localStorage.getItem("user");
    console.log("Changement de rôle vers:", newRole);
    if (userString) {
      try {
        const user = JSON.parse(userString);
        user.role = newRole;
        localStorage.setItem("user", JSON.stringify(user));
        setRole(newRole);
      } catch (error) {
        console.error("Erreur lors du changement de rôle:", error);
      }
    }
  };

  return (
    <AnimatePresence mode="wait">
      <Routes location={location} key={location.pathname}>
        <Route path="/" element={<Navigate to="/login" replace />} />

        {/* Authentification */}
        <Route
          path="/login"
          element={
            <PageWrapper>
              <LoginPage />
            </PageWrapper>
          }
        />
        <Route
          path="/signup"
          element={
            <PageWrapper>
              <Signup />
            </PageWrapper>
          }
        />

        {/* Paramètre */}
        <Route
          path="/parametre"
          element={
            <PageWrapper>
              <Parametre />
            </PageWrapper>
          }
        />

        {/* Pages principales */}
        <Route
          path="/alerts"
          element={
            <PageWrapper>
              <Alerts />
            </PageWrapper>
          }
        />
        <Route
          path="/modules"
          element={
            <PageWrapper>
              <ModuleManagement role={role} onRoleChange={handleRoleChange} />
            </PageWrapper>
          }
        />
        
        {/* CORRECTION: Seuls les chefs de département peuvent accéder aux commentaires */}
        <Route
          path="/commentaires"
          element={
            role === "chefDepartement" ? (
              <PageWrapper>
                <Commentaires />
              </PageWrapper>
            ) : (
              <Navigate to="/modules" replace />
            )
          }
        />
        
        <Route
          path="/wishlist"
          element={
            <PageWrapper>
              <WishList />
            </PageWrapper>
          }
        />
        <Route
          path="/profil"
          element={
            <PageWrapper>
              <Profil />
            </PageWrapper>
          }
        />
        <Route
          path="/enseignants"
          element={
            <PageWrapper>
              <TeacherTableManagment role={role} />
            </PageWrapper>
          }
        />
        <Route
          path="/organigramme"
          element={
            <PageWrapper>
              <OrganigrammePage />
            </PageWrapper>
          }
        />

        {/* Dashboards */}
        <Route
          path="/dashboardtec"
          element={
            <PageWrapper>
              <Dashboardtec />
            </PageWrapper>
          }
        />
        <Route
          path="/dashboardorga"
          element={
            <PageWrapper>
              <Dashboardorga userRole={role} />
            </PageWrapper>
          }
        />
        <Route
          path="/codes"
          element={
            role === "chefDepartement" ? (
              <PageWrapper>
                <CodesPage />
              </PageWrapper>
            ) : (
              <Navigate to="/" replace />
            )
          }
        />

        {/* 404 fallback */}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </AnimatePresence>
  );
};

export default App;