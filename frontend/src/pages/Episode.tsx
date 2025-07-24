import React from 'react';
import { useParams } from 'react-router-dom';
import { useQuery } from 'react-query';
import axios from 'axios';
import { Play, Lock, Unlock } from 'lucide-react';

interface Episode {
  id: string;
  title: string;
  description: string;
  audioUrl: string;
  duration: number;
  coinPrice: number;
  isOwned: boolean;
  canUnlock: boolean;
}

const Episode: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  const { data: episode, isLoading, error } = useQuery<Episode>(
    ['episode', id],
    async () => {
      const response = await axios.get(`/api/v1/episodes/${id}`);
      return response.data;
    }
  );

  const handleUnlock = async () => {
    try {
      await axios.post(`/api/v1/episodes/${id}/unlock`);
      // Refetch episode data
      window.location.reload();
    } catch (error) {
      console.error('Failed to unlock episode:', error);
    }
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  if (error || !episode) {
    return (
      <div className="text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Error loading episode</h2>
        <p className="text-gray-600">Please try again later.</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <div className="bg-white rounded-lg shadow-lg p-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-4">{episode.title}</h1>
        <p className="text-gray-600 mb-6">{episode.description}</p>
        
        <div className="mb-6">
          <p className="text-gray-500">
            Duration: {Math.floor(episode.duration / 60)}:{String(episode.duration % 60).padStart(2, '0')}
          </p>
        </div>

        {episode.isOwned ? (
          <div className="bg-green-50 border border-green-200 rounded-lg p-6">
            <div className="flex items-center space-x-3 mb-4">
              <Play className="h-6 w-6 text-green-600" />
              <h3 className="text-lg font-semibold text-green-800">Episode Unlocked</h3>
            </div>
            <p className="text-green-700 mb-4">You can now listen to this episode.</p>
            <audio controls className="w-full">
              <source src={episode.audioUrl} type="audio/mpeg" />
              Your browser does not support the audio element.
            </audio>
          </div>
        ) : (
          <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
            <div className="flex items-center space-x-3 mb-4">
              <Lock className="h-6 w-6 text-yellow-600" />
              <h3 className="text-lg font-semibold text-yellow-800">Episode Locked</h3>
            </div>
            <p className="text-yellow-700 mb-4">
              Unlock this episode for {episode.coinPrice} coins to start listening.
            </p>
            <button
              onClick={handleUnlock}
              disabled={!episode.canUnlock}
              className={`flex items-center space-x-2 px-6 py-3 rounded-lg font-semibold transition-colors ${
                episode.canUnlock
                  ? 'bg-indigo-600 text-white hover:bg-indigo-700'
                  : 'bg-gray-300 text-gray-500 cursor-not-allowed'
              }`}
            >
              <Unlock className="h-5 w-5" />
              <span>
                {episode.canUnlock ? `Unlock for ${episode.coinPrice} coins` : 'Insufficient coins'}
              </span>
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Episode; 