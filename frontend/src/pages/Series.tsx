import React from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from 'react-query';
import axios from 'axios';
import { Play, Star } from 'lucide-react';

interface Series {
  id: string;
  title: string;
  description: string;
  coverImage: string;
  author: string;
  category: string;
  isPremium: boolean;
  totalEpisodes: number;
}

const Series: React.FC = () => {
  const { data: series, isLoading, error } = useQuery<Series[]>(
    'series',
    async () => {
      const response = await axios.get('/api/v1/series');
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

  if (error) {
    return (
      <div className="text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Error loading series</h2>
        <p className="text-gray-600">Please try again later.</p>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-4">Audio Series</h1>
        <p className="text-gray-600">Discover amazing audio stories and unlock episodes with coins</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        {series?.map((item) => (
          <div key={item.id} className="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-shadow">
            <div className="relative">
              <img
                src={item.coverImage}
                alt={item.title}
                className="w-full h-48 object-cover"
              />
              {item.isPremium && (
                <div className="absolute top-2 right-2 bg-yellow-500 text-white px-2 py-1 rounded-full text-xs font-semibold flex items-center">
                  <Star className="h-3 w-3 mr-1" />
                  Premium
                </div>
              )}
            </div>
            <div className="p-6">
              <h3 className="text-xl font-semibold text-gray-900 mb-2">{item.title}</h3>
              <p className="text-gray-600 mb-4 line-clamp-2">{item.description}</p>
              <div className="flex items-center justify-between">
                <div className="text-sm text-gray-500">
                  <p>By {item.author}</p>
                  <p>{item.totalEpisodes} episodes</p>
                </div>
                <Link
                  to={`/series/${item.id}`}
                  className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition-colors flex items-center"
                >
                  <Play className="h-4 w-4 mr-2" />
                  Listen
                </Link>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Series; 