import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import './index.css';
import TeamsList from './teams/list.js';
import TeamMatches from './team/matches.js';
import NotFound from './404.js';

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render(
  <React.StrictMode>
    <div id="container">
      <h1>Team 846 Scouting - Chezy Champs</h1>

      <BrowserRouter>
        <Routes>
          <Route path="/" element={<TeamsList />} />
          <Route path="team/:teamNumber" element={<TeamMatches />} />

          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </div>
  </React.StrictMode>
);
