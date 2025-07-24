import React from 'react';
import { useQuery } from 'react-query';
import axios from 'axios';
import { useAuth } from '../contexts/AuthContext';
import { User, Coins, ShoppingBag } from 'lucide-react';

interface Purchase {
  id: string;
  type: string;
  amount: number;
  status: string;
  createdAt: string;
}

const Profile: React.FC = () => {
  const { user } = useAuth();

  const { data: purchases, isLoading } = useQuery<Purchase[]>(
    'purchases',
    async () => {
      const response = await axios.get('/api/v1/user/purchases');
      return response.data;
    },
    {
      enabled: !!user,
    }
  );

  if (!user) {
    return (
      <div className="text-center py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">Please log in</h2>
        <p className="text-gray-600">You need to be logged in to view your profile.</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto">
      <div className="bg-white rounded-lg shadow-lg p-8">
        <div className="flex items-center space-x-4 mb-8">
          <div className="w-16 h-16 bg-indigo-100 rounded-full flex items-center justify-center">
            <User className="h-8 w-8 text-indigo-600" />
          </div>
          <div>
            <h1 className="text-3xl font-bold text-gray-900">
              {user.firstName} {user.lastName}
            </h1>
            <p className="text-gray-600">{user.email}</p>
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          {/* Coin Balance */}
          <div className="bg-gradient-to-r from-yellow-400 to-yellow-600 rounded-lg p-6 text-white">
            <div className="flex items-center space-x-3 mb-4">
              <Coins className="h-8 w-8" />
              <h2 className="text-2xl font-bold">Coin Balance</h2>
            </div>
            <p className="text-4xl font-bold">{user.coinBalance}</p>
            <p className="text-yellow-100">Available coins</p>
          </div>

          {/* Account Info */}
          <div className="bg-gray-50 rounded-lg p-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Account Information</h2>
            <div className="space-y-2">
              <p><span className="font-medium">Role:</span> {user.role}</p>
              <p><span className="font-medium">Member since:</span> {new Date().toLocaleDateString()}</p>
            </div>
          </div>
        </div>

        {/* Purchase History */}
        <div className="mt-8">
          <div className="flex items-center space-x-3 mb-6">
            <ShoppingBag className="h-6 w-6 text-gray-600" />
            <h2 className="text-2xl font-bold text-gray-900">Purchase History</h2>
          </div>
          
          {isLoading ? (
            <div className="flex items-center justify-center py-8">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
            </div>
          ) : purchases && purchases.length > 0 ? (
            <div className="space-y-4">
              {purchases.map((purchase) => (
                <div key={purchase.id} className="flex items-center justify-between p-4 border border-gray-200 rounded-lg">
                  <div>
                    <p className="font-semibold text-gray-900">{purchase.type}</p>
                    <p className="text-gray-600 text-sm">
                      {new Date(purchase.createdAt).toLocaleDateString()}
                    </p>
                  </div>
                  <div className="text-right">
                    <p className="font-semibold text-gray-900">{purchase.amount} coins</p>
                    <span className={`text-sm px-2 py-1 rounded-full ${
                      purchase.status === 'completed' 
                        ? 'bg-green-100 text-green-800' 
                        : 'bg-yellow-100 text-yellow-800'
                    }`}>
                      {purchase.status}
                    </span>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="text-center py-8 text-gray-500">
              <p>No purchase history yet.</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default Profile; 