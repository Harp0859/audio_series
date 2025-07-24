import React from 'react';
import { useParams } from 'react-router-dom';
import { useQuery } from 'react-query';
import axios from 'axios';
import { Play, Lock, Star } from 'lucide-react';

interface Episode {
  id: string;
  title: string;
  description: string;
  episodeNumber: number;
  duration: number;
  coinPrice: number;
  isLocked: boolean;
}

interface SeriesWithEpisodes {
  series: {
    id: string;
    title: string;
    description: string;
    coverImage: string;
    author: string;
    category: string;
    isPremium: boolean;
    totalEpisodes: number;
  };
  episodes: Episode[];
}

const SeriesDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  const { data: seriesData, isLoading, error } = useQuery<SeriesWithEpisodes>(
    ['series', id],
    async () => {
      const response = await axios.get(`/api/v1/series/${id}`);
      return response.data;
    }
  );

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  if (error || !seriesData) {
    return (
      <div className="text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Error loading series</h2>
        <p className="text-gray-600">Please try again later.</p>
      </div>
    );
  }

  const { series, episodes } = seriesData;

  return (
    <div className="max-w-7xl mx-auto">
      <div className="mb-8">
        <div className="flex items-center space-x-4">
          <img
            src={series.coverImage}
            alt={series.title}
            className="w-32 h-32 object-cover rounded-lg"
          />
          <div>
            <h1 className="text-3xl font-bold text-gray-900 mb-2">{series.title}</h1>
            <p className="text-gray-600 mb-2">By {series.author}</p>
            <p className="text-gray-600 mb-4">{series.description}</p>
            {series.isPremium && (
              <div className="inline-flex items-center bg-yellow-500 text-white px-3 py-1 rounded-full text-sm font-semibold">
                <Star className="h-4 w-4 mr-1" />
                Premium Series
              </div>
            )}
          </div>
        </div>
      </div>

      <div className="bg-white rounded-lg shadow-lg p-6">
        <h2 className="text-2xl font-bold text-gray-900 mb-6">Episodes</h2>
        <div className="space-y-4">
          {episodes.map((episode) => (
            <div key={episode.id} className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
              <div className="flex items-center space-x-4">
                <div className="w-12 h-12 bg-indigo-100 rounded-full flex items-center justify-center">
                  <span className="text-indigo-600 font-semibold">{episode.episodeNumber}</span>
                </div>
                <div>
                  <h3 className="font-semibold text-gray-900">{episode.title}</h3>
                  <p className="text-gray-600 text-sm">{episode.description}</p>
                  <p className="text-gray-500 text-sm">
                    {Math.floor(episode.duration / 60)}:{String(episode.duration % 60).padStart(2, '0')}
                  </p>
                </div>
              </div>
              <div className="flex items-center space-x-2">
                {episode.isLocked ? (
                  <div className="flex items-center space-x-2">
                    <Lock className="h-5 w-5 text-gray-400" />
                    <span className="text-gray-500">{episode.coinPrice} coins</span>
                  </div>
                ) : (
                  <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors flex items-center">
                    <Play className="h-4 w-4 mr-2" />
                    Play
                  </button>
                )}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default SeriesDetail; 