import React from 'react';
import { Link } from 'react-router-dom';
import { useQuery } from 'react-query';
import axios from 'axios';
import { Play, Lock, Star } from 'lucide-react';

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

const Home: React.FC = () => {
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
      {/* Hero Section */}
      <div className="text-center py-12 bg-gradient-to-r from-indigo-600 to-purple-600 text-white rounded-lg mb-12">
        <h1 className="text-4xl font-bold mb-4">Welcome to AudioSeries</h1>
        <p className="text-xl mb-8">Discover amazing audio stories and unlock episodes with coins</p>
        <div className="flex justify-center space-x-4">
          <Link
            to="/series"
            className="bg-white text-indigo-600 px-6 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors"
          >
            Browse Series
          </Link>
          <Link
            to="/register"
            className="bg-transparent border-2 border-white text-white px-6 py-3 rounded-lg font-semibold hover:bg-white hover:text-indigo-600 transition-colors"
          >
            Get Started
          </Link>
        </div>
      </div>

      {/* Featured Series */}
      <div className="mb-12">
        <h2 className="text-3xl font-bold text-gray-900 mb-8">Featured Series</h2>
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

      {/* How it Works */}
      <div className="bg-gray-50 rounded-lg p-8 mb-12">
        <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center">How It Works</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div className="text-center">
            <div className="bg-indigo-600 text-white rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
              1
            </div>
            <h3 className="text-xl font-semibold mb-2">Browse Series</h3>
            <p className="text-gray-600">Explore our collection of audio series across different genres</p>
          </div>
          <div className="text-center">
            <div className="bg-indigo-600 text-white rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
              2
            </div>
            <h3 className="text-xl font-semibold mb-2">Unlock Episodes</h3>
            <p className="text-gray-600">Use your coins to unlock individual episodes or entire series</p>
          </div>
          <div className="text-center">
            <div className="bg-indigo-600 text-white rounded-full w-16 h-16 flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
              3
            </div>
            <h3 className="text-xl font-semibold mb-2">Enjoy Listening</h3>
            <p className="text-gray-600">Stream your unlocked episodes anytime, anywhere</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home; 