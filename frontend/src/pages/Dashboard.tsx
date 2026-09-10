import { useEffect } from "react";
import { apiRequest } from "../services/api";

const Dashboard = () => {
  // function to check if backend is healthy
  const checkBackend = async () => {
    try {
      const data = await apiRequest("/v1/health");

      console.log("API response : ", data);
    } catch (err) {
      console.log("Error is : ", err);
    }
  };

  useEffect(() => {
    checkBackend();
  }, []);

  return <div>Dashboard</div>;
};

export default Dashboard;
